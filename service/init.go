package service

import (
	"../model"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	// 下划线作用是 帮忙 运行 包里面的 init函数
	_ "github.com/go-sql-driver/mysql"
)

// 申明一个 指针型的 操作引擎
var DbEngine *xorm.Engine

func init()  {
	driverName := "mysql"
	dataSourceName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	var err error = errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, dataSourceName)
	if nil != err && err.Error() != "" {
		// 打印报错信息 并退出
		log.Fatal(err.Error())
	}

	// 是否 打印 SQL
	DbEngine.ShowSQL(true)
	// 设置数据库打开的连接数
	DbEngine.SetMaxOpenConns(2)

	// 是否开启自动建表
	DbEngine.Sync2(new(model.User))
	fmt.Printf("init data base access engine ok!")
}
