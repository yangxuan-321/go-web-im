package util

import (
	"encoding/json"
	"net/http"
)


// 配置 json , 免得 返回的json 首字母 大写
// omitempty 对 nil 不显示
type Result struct {
	Code int			`json:"code"`
	Msg string			`json:"msg"`
	Data interface{}	`json:"data,omitempty"`
}

func RespSuccess(writer http.ResponseWriter, code int, data interface{}, msg string)  {
	Resp(writer, 0, data, msg)
}

func RespError(writer http.ResponseWriter, code int, data interface{}, msg string)  {
	Resp(writer, -1, nil, msg)
}


func Resp(writer http.ResponseWriter, code int, data interface{}, msg string)  {
	// 设置header为json, header默认的是text/html, 所以特别指出返回的
	// 设置为 application/json
	writer.Header().Set("Content-Type", "application/json")
	// 设置200状态
	writer.WriteHeader(http.StatusOK)

	// 定义返回类型
	var result = Result{
		Code:code,
		Msg:msg,
		Data:data,
	}

	resultJson, _ := json.Marshal(result)

	// 输出
	writer.Write([]byte(resultJson))
}