package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

//订制配置文件解析载体
type Config struct{
	Liexin_databases *LiexinMysqlConfig
	Rabbitmq_ichunt *RabbitmqIchunt
	Crm_domain *SendMail
	Ding_msg *Ding
	Api_domain *ApiDomain
	Mongodb_databases *MongodbDatabases
	LogSaveConfig *LogSaveConfigs
}


type LiexinMysqlConfig struct{
	Dns string `toml:"dns"`
}

type RabbitmqIchunt struct {
	QueueName string `toml:"queue_name"`
	RoutingKey string `toml:"routing_key"`
	Exchange string `toml:"exchange"`
	Type string `toml:"type"`
	Dns string `toml:"dns"`
}

type MongodbDatabases struct {
	Dns string `toml:"dns"`
	Databases string `toml:"databases"`
	Collection string `toml:"collection"`
}

type  LogSaveConfigs struct {
	Time int `toml:"time"`
	Length int `toml:"length"`
}

type SendMail struct{
	SendMailUrl string `toml:"send_mail"`
}

type Ding struct {
	Webhook string  `toml:"webhook"`
	JingDiao string `toml:"jingDiao"`
}

type ApiDomain struct {
	ApiUrl string `toml:"api_url"`
}

var Configs *Config =new (Config)

func Init(ConfigDir string){
	//fmt.Println(ConfigDir+"config/config.toml")
	var err error
	_, err = toml.DecodeFile(ConfigDir+"config.toml",Configs)
	_, err = toml.DecodeFile(ConfigDir+"db.toml",Configs)
	if err!=nil{
		fmt.Println(err)
	}
	//fmt.Printf("%+v",Configs.Liexin_databases)
	//fmt.Printf("%+v",Configs.Rabbitmq_ichunt)
	//fmt.Printf("%+v",Configs.Crm_domain)
	//fmt.Printf("%+v",Configs.Ding_msg)
	//
	//fmt.Printf("%+v",Configs.Crm_domain)
	//fmt.Printf("%+v",Configs.Rabbitmq_ichunt)
	//fmt.Printf("%+v",Configs.Mongodb_databases)
	//fmt.Printf("%+v",Configs.LogSaveConfig)
}
