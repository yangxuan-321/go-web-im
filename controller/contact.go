package controller

import (
	"../args/contact"
	"../args/user"
	"../service"
	"../util"
	"net/http"
)

var contactService *service.ContactService = &service.ContactService{}

func Addfriend(writer http.ResponseWriter, request *http.Request) {
	// 通过反射工具将 前端 传过来的 参数 绑定至 对象上
	var arg args_contact.ContactArg
	util.Bind(request, &arg)

	// 调用service 进行 添加好友 逻辑
	err := contactService.Addfriend(arg.Userid, arg.Dstmobile)

	if err != nil {
		util.RespError(writer, "添加好友失败:"+err.Error())
	} else {
		util.RespSuccess(writer, nil, "添加好友成功")
	}
}

func LoadFriendAndGroup(writer http.ResponseWriter, request *http.Request) {
	// 先判断是否登录
	// 通过反射工具将 前端 传过来的 参数 绑定至 对象上
	var arg *args_user.UserTokenArg = &args_user.UserTokenArg{}
	util.Bind(request, arg)

	vos, err := contactService.LoadFriendAndGroup(arg)

	if nil != err {
		util.RespError(writer, "加载好友列表失败:"+err.Error())
	} else {
		util.RespList(writer, 0, vos, len(vos))
	}
}
