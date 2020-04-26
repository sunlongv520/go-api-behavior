package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/ichunt2019/logger"
)

//2018/3/26 0:01.383 DEBUG logDebug.go:29 this is a debug log
//2006-01-02 15:04:05.999
type FileLogger struct {
	level         int
	logPath       string
	logName       string
	file          *os.File
	warnFile      *os.File
	LogDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}
type LogData struct {
	Message      string
}

// 队列参数
type DataParams struct {
	InterfaceType string `json:"interface_type"`
	AccessUrl string `json:"access_url"`
	ErrMsg string `json:"err_msg"`
	ErrCode string `json:"err_code"`
	Uid string `json:"uid"`
	UserName string `json:"user_name"`
	UserIp string `json:"user_ip"`
	Remakr string `json:"remark"`
	CreateTime int64 `json:"create_time"`
	CreateTimeStr string `json:"create_time_str"`
}

var MsgChan chan *DataParams
var autoCommitChan chan *DataParams
var MsgParams *DataParams

// 日志批次
type LogBatch struct {
	Logs []*DataParams	// 多条日志
}

func NewFileLogger(LogReportDir string) (log FileLogger, err error) {
	logPath := LogReportDir

	logName := time.Now().Format("2006-01-02")

	logSplitSize, err := strconv.ParseInt("104857600", 10, 64)
	if err != nil {
		logSplitSize = 104857600
	}
	logChanSize := "50000"

	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	log = FileLogger{
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *LogData, chanSize),
		logSplitSize:  logSplitSize,
	}

	return
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string)  error  {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (f *FileLogger) Init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	createFile(f.logPath)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filename, err))
	}

	f.file = file
}

func (f *FileLogger) splitFileHour(warnFile bool) {
	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}

	f.lastSplitHour = hour
	var backupFilename string
	var filename string

	backupFilename = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d",
		f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
	filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)

	file := f.file

	file.Close()
	os.Rename(filename, backupFilename)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	f.file = file
}

func (f *FileLogger) splitFileSize() {

	file := f.file


	statInfo, err := file.Stat()
	if err != nil {
		return
	}

	fileSize := statInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}

	var backupFilename string
	var filename string

	now := time.Now()
	backupFilename = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d%02d",
		f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)

	file.Close()
	os.Rename(filename, backupFilename)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	f.file = file
}

func (f *FileLogger) checkSplitFile() {

	f.splitFileSize()
}

func (f *FileLogger) WriteLog(logBatch *LogBatch) {
	for _,logData := range logBatch.Logs {
		fmt.Println("开始写入日志.....")
		fmt.Println(logData)
		var file *os.File = f.file
		f.checkSplitFile()
		data, err := json.Marshal(logData)
		if err != nil{
			logger.Error("%s",err)
			continue
		}
		fmt.Fprintf(file, "%s\n", string(data))
	}
}



func (f *FileLogger) Close() {
	f.file.Close()
}
