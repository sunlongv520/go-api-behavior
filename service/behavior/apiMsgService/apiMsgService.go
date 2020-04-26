package apiMsgService

import (
	"encoding/json"
	"time"
)

// 队列参数
type DataParams struct {
	InterfaceType string `json:"interface_type"`
	AccessUrl string `json:"access_url"`
	RequestParams string `json:"request_params"`
	ErrMsg string `json:"err_msg"`
	ErrCode string `json:"err_code"`
	Uid string `json:"uid"`
	UserName string `json:"user_name"`
	UserIp string `json:"user_ip"`
	Remakr string `json:"remark"`
	CreateTime int64 `json:"create_time"`
	CreateTimeStr string `json:"create_time_str"`
}

var MsgParams *DataParams


//写日志到管道
func  SaveMsgToLogChan(msg string) (err error) {
	err = json.Unmarshal([]byte(msg),&MsgParams)
	MsgParams.CreateTime = time.Now().Unix()
	MsgParams.CreateTimeStr = time.Now().Format("2006-01-02 15:04:05")
	G_logSink.Append(MsgParams)
	return err
}
