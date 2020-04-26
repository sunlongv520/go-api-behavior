package ding

import (
	_"fmt"
	"go-api-behavior/util"
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
	"github.com/ichunt2019/logger"
)

type Msg struct {
  Msgtype string `json:"msgtype"`
  Text Contents `json:"text"`
  At Ats `json:"at"`
}

type  Contents struct {
  Content string `json:"content"`
}

type Ats struct {
  AtMobiles []string `json:"atMobiles"`
  IsAtAll bool `json:"isAtAll"`
}

// json 返回值
type JosnResp struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

func Send(textMsg string, mobiles []string, isAtAll bool) (jsonStr string) {
	var msg Msg
	msg = Msg{Msgtype:"text"}
	msg.Text.Content = "监控告警: " + textMsg // 固定标签 + 文本
	
	msg.At.AtMobiles = mobiles
	msg.At.IsAtAll = isAtAll; 

	content, _ := json.Marshal(msg)

	// content := `{
	// 	"msgtype": "text",
	// 	"text": {
	// 		"content": "`+ msg + `"
	// 	}
	// }`
	
	ding_url := util.Configs.Ding_msg.Webhook // 钉钉请求接口

	client := &http.Client{}

	req, _ := http.NewRequest("POST", ding_url, strings.NewReader(string(content)))

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req);

	defer resp.Body.Close()

    if err != nil {
        logger.Info(err.Error())
    }

    body, _ := ioutil.ReadAll(resp.Body) // 获取接口返回数据

	res := JosnResp{}

	if err1 := json.Unmarshal([]byte(body), &res); err1 != nil { // 将接口数据写入res
		logger.Info(err1.Error())
    }

    result, _ := json.Marshal(res)

    return string(result)
}