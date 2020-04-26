package apiMsgService

import (
	"go-api-behavior/util"

	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


type LogBatch struct {
	Logs []interface{}	// 多条日志
}


// mongodb存储日志
type LogSink struct {
	client *mongo.Client
	logCollection *mongo.Collection
	logChan chan *DataParams
	autoCommitChan chan *LogBatch
}

var (
	// 单例
	G_logSink *LogSink
)



// 批量写入日志
func (logSink *LogSink) saveLogs(batch *LogBatch) {
	logSink.logCollection.InsertMany(context.TODO(), batch.Logs)
}

// 日志存储协程
func (logSink *LogSink) writeLoop() {
	//select{
	//case log := <- logSink.logChan:
	//	fmt.Println("从管道中读取日志")
	//	fmt.Println(log)
	//}
	var (
		log *DataParams
		logBatch *LogBatch // 当前的批次
		//commitTimer *time.Timer
		//timeoutBatch *LogBatch // 超时批次
	)

	sleepTime :=  time.Duration(int64(util.Configs.LogSaveConfig.Time))
	t := time.NewTimer(time.Second * sleepTime)
	defer t.Stop()
	for {
		select {
		case log = <- logSink.logChan:
			//fmt.Println("新任务")
			if logBatch == nil {
				logBatch = &LogBatch{}
			}

			// 把新日志追加到批次中
			logBatch.Logs = append(logBatch.Logs, *log)
			// 如果批次满了, 就立即发送
			if len(logBatch.Logs) >= int(util.Configs.LogSaveConfig.Length) {
				// 发送日志
				logSink.saveLogs(logBatch)
				// 清空logBatch
				logBatch = nil
			}
		case <- t.C:
			if logBatch == nil {
				logBatch = &LogBatch{}
			}
			//fmt.Println("超时到期了")
			//fmt.Println(len(logBatch.Logs))
			logSink.saveLogs(logBatch)
			logBatch = nil
			t.Reset(time.Second * sleepTime)
		}


	}
}

func InitLogSink() (err error) {
	var (
		client *mongo.Client
	)

	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI(util.Configs.Mongodb_databases.Dns)
	if client, err = mongo.Connect(
		context.TODO(),clientOptions); err != nil {
		return
	}

	//   选择db和collection
	G_logSink = &LogSink{
		client: client,
		logCollection: client.Database(util.Configs.Mongodb_databases.Databases).Collection(util.Configs.Mongodb_databases.Collection),
		logChan: make(chan *DataParams, 1000),
		//autoCommitChan: make(chan *common.LogBatch, 1000),
	}

	// 启动一个mongodb处理协程
	go G_logSink.writeLoop()
	return
}

// 发送日志
func (logSink *LogSink) Append(jobLog *DataParams) {
	select {
	case logSink.logChan <- jobLog:
	default:
		// 队列满了就丢弃
	}
}