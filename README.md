## 本地运行
**go run .\service\behavior\main.go  -configDir=./config/  -logDir=./log/ -LogReportDir=./LogReport/**


## 生产
** go build -o {path}/cmd/master  {path}/service\behavior\main.go**


## 测试json
    {
    	"interface_type":"1",
    	"access_url":"api.ichun.com/v3/login",
    	"err_msg":"响应超时",
    	"err_code":"40004",
    	"request_params":"tttttttttt",
    	"uid":"16678",
    	"user_name":"张三",
    	"user_ip":"153.2.15.1",
    	"remark":"a=1&b=2&c=3",
    	"create_time":1581582711,
    	"create_time_str":"2020-02-14 16:40"
    }

## 配置修改
复制config目录下的
cp  ./config/db.toml.demo   ./config/db.toml
修改相关配置即可