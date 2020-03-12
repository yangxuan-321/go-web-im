package controller

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

// 聊天
// 请求url 类似于 ws://127.0.0.1/chat?id=1&token=xxx
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1.获取请求参数 id和token
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	token := query.Get("token")

	// 2.校验 token
	isValida := checkToken(id, token)

	// 3.接入连接
	conn, err := (&websocket.Upgrader{
		// 检测请求来源
		// 这个 检测 是否成功 是与 token 紧密绑定的。token 有效 就成功
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

//检测是否有效
func checkToken(userId int64, token string) bool {
	//从数据库里面查询并比对
	// 此处后期 需要存入redis
	user, err := userService.FindUserById(userId)
	if err != nil {
		return false
	}
	return user.Token == token
}
