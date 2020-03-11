package controller

import (
	"../model"
	"../service"
	"../util"
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
		util.RespError(writer, "密码不正确")
	} else {
		util.RespSuccess(writer, user, "密码正确")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	// 1.如何获得参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	//niclname := fmt.Sprint("user%06d", rand.Int31())
	nickname := util.RandNickname()
	avatar := "/asset/images/" + util.RandAvatar()
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if nil != err {
		util.RespError(writer, "注册失败:"+err.Error())
	} else {
		util.RespSuccess(writer, user, "注册成功")
	}
}
