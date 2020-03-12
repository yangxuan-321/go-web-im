package main

import (
	"./controller"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/gpmgo/gopm/modules/log"
	"html/template"
	"net/http"
)

func RegisterView() {

	// 1.解析template 得到一个模板的指针
	// 满足 相应格式 的 都可以 被解析出来
	tpl, err := template.ParseGlob("view/**/*.html")
	if nil != err {
		// 打印并直接退出
		log.Fatal(err.Error())
	}

	// 循环拿到 template
	for _, v := range tpl.Templates() {
		// 这个地方的 Name， 就是 通过在 模板文件[html文件] 中 加入 {{define "xxx/xxx/xx.shtml"}} 拿到
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

// 申明一个 指针型的 操作引擎
var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	dataSourceName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	DbEngine, err := xorm.NewEngine(driverName, dataSourceName)
	if nil != err {
		// 打印报错信息 并退出
		log.Fatal(err.Error())
	}

	// 是否 打印 SQL
	DbEngine.ShowSQL(true)
	// 设置数据库打开的连接数
	DbEngine.SetMaxOpenConns(2)

	// 是否开启自动建表
	//DbEngine.Sync2(new(User))
	fmt.Printf("init data base access engine ok!")
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriendAndGroup)
	http.HandleFunc("/contact/createcommunity", controller.CreateGroup)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)

	http.HandleFunc("/chat", controller.Chat)

	// 提供静态资源 目录支持interceptors
	// 第一个参数 / 代表 访问路径
	// 第二个参数 . 代表 当前目录
	//http.Handle("/", http.FileServer(http.Dir(".")))

	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// 对渲染模板的访问 注册
	RegisterView()

	// 启动web服务器
	http.ListenAndServe(":8080", nil)
}
