package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResultList struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

// 配置 json , 免得 返回的json 首字母 大写
// omitempty 对 nil 不显示
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespSuccess(writer http.ResponseWriter, data interface{}, msg string) {
	Resp(writer, 0, data, msg)
}

func RespError(writer http.ResponseWriter, msg string) {
	Resp(writer, -1, nil, msg)
}

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	// 设置header为json, header默认的是text/html, 所以特别指出返回的
	// 设置为 application/json
	writer.Header().Set("Content-Type", "application/json")
	// 设置200状态
	writer.WriteHeader(http.StatusOK)

	// 定义返回类型
	var result = Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	resultJson, _ := json.Marshal(result)

	// 输出
	writer.Write([]byte(resultJson))
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20
	h := ResultList{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
