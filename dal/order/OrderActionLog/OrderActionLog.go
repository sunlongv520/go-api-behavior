package OrderActionLog

import (
	"database/sql"
	"fmt"
	"github.com/ichunt2019/logger"
	"go-api-behavior/dal/db"
	"go-api-behavior/util"
	"log"
	"time"
)

type ActionLog struct {
    LogId   int    `db:"log_id"`
    OrderId int `db:"order_id"`
    OperatorId  int `db:"operator_id"`
    OperatorType int `db:"operator_type"`
    Event string `db:"event"`
    Ip string `db:"ip"`
    CreateTime int `db:"create_time"`
}

func initDb(dns string) (err error) {
	err = db.Init(dns)
	if err != nil {
		return
	}

	return
}

func AddLog() (err error) {
	initDb(util.Configs.Liexin_databases.Dns) //初始化db

	// QueryRow()
	// Query()
	Insert(3340, 1000, 2, "go test")

	return
}

// 单行查询
func QueryRow() {
	var actionLog ActionLog

	err1 := db.DB.Get(&actionLog, "select * from lie_order_action_log where order_id = 3340")

	if err1 == sql.ErrNoRows {
        log.Printf("not found data of the id:%d", 1)
    }

    if err1 != nil {
        panic(err1)
    }

    fmt.Printf("actionLog: %v\n", actionLog)
}

// 多行查询
func Query() {
	var actionLog2 []*ActionLog
    err2 := db.DB.Select(&actionLog2, "select * from lie_order_action_log where order_id = 3340")

    if err2 == sql.ErrNoRows {
        log.Printf("not found data of the id:%d", 1)
    }

    if err2 != nil {
        panic(err2)
    }

    for _, v := range actionLog2 {
        fmt.Println(v)
    }
}

func Insert(order_id int, operator_id int, operator_type int, event string) (err error) {
	create_time := time.Now().Unix()

    _, err1 := db.DB.Exec("insert into lie_order_action_log (order_id, operator_id, operator_type, event, create_time) values (?, ?, ?, ?, ?)", order_id, operator_id, operator_type, event, create_time)

    if err1 != nil {
        logger.Fatal("数据库操作失败")
        return err1
    }

    //id, err := result.LastInsertId()
    //if err != nil {
    //    panic(err)
    //}
	//
    //affected, err := result.RowsAffected()
    //if err != nil {
    //    panic(err)
    //}
	//
    //fmt.Printf("last insert id:%d affect rows:%d\n", id, affected)

    return err
}