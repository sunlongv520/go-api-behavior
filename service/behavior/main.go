package main

import (
	"flag"
	"github.com/ichunt2019/go-rabbitmq/utils/rabbitmq"
	"github.com/ichunt2019/logger"
	"go-api-behavior/service/behavior/apiMsgService"
	"go-api-behavior/util"
)

type RecvPro struct {

}

//// 实现消费者 消费消息失败 自动进入延时尝试  尝试3次之后入库db
func (t *RecvPro) Consumer(dataByte []byte) error {
	//fmt.Println(string(dataByte))
	err :=  apiMsgService.SaveMsgToLogChan(string(dataByte))
	if err != nil{
		logger.Error("%s",err)
		return err
	}
	////return errors.New("顶顶顶顶")
	return nil
}

//消息已经消费3次 失败了 请进行处理
func (t *RecvPro) FailAction(dataByte []byte) error {
	logger.Error("任务处理失败了，我要进入db日志库了")
	logger.Error("任务处理失败了，发送钉钉消息通知主人")
	logger.Error(string(dataByte))
	return nil
}


var ConfigDir string
var LogDir string
var LogReportDir string

// 解析命令行参数
func initArgs() {
	// worker -config ./worker.json
	// worker -h
	flag.StringVar(&ConfigDir, "configDir", "", "配置文件")
	flag.StringVar(&LogDir, "logDir", "", "日志目录")
	flag.StringVar(&LogReportDir, "LogReportDir", "", "用户行为日志目录")
	flag.Parse()
}


func main() {

	initArgs()

	//初始化配置文件
	util.Init(ConfigDir)


	//
	logConfig := make(map[string]string)
	logConfig["log_path"] = LogDir+"api/behavior"
	logConfig["log_chan_size"] = "1000"
	logger.InitLogger("file",logConfig)
	logger.Init()
	//初始化db
	//initDb(util.Configs.Liexin_databases.Dns)

	t := &RecvPro{}



	_ = apiMsgService.InitLogSink()
	



	rabbitmq.Recv(rabbitmq.QueueExchange{
		util.Configs.Rabbitmq_ichunt.QueueName,
		util.Configs.Rabbitmq_ichunt.RoutingKey,
		util.Configs.Rabbitmq_ichunt.Exchange,
		util.Configs.Rabbitmq_ichunt.Type,
		util.Configs.Rabbitmq_ichunt.Dns,
	},t,3)




}