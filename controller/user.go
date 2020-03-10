package controller

import (
	"../model"
	"../service"
	"../util"
	"fmt"
	"math/rand"
	"net/http"
)

var userService *service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	// 1.如何获得参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	user, err := userService.Login(mobile, passwd)


	if nil != err {
		util.RespError(writer, -1, nil, "密码不正确")
	}else {
		util.RespSuccess(writer, 0, user, "密码正确")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request)  {
	// 1.如何获得参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	niclname := fmt.Sprint("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, niclname, avatar, sex)
	if nil != err {
		util.RespError(writer, -1, nil, err.Error())
	}else {
		util.RespSuccess(writer, 0, user, "")
	}
}
