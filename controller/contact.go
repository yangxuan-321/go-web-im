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
	var arg args_contact.ContactUserArg
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

// 创建群聊
func CreateGroup(writer http.ResponseWriter, request *http.Request) {
	// 通过反射工具将 前端 传过来的 参数 绑定至 对象上
	var arg args_contact.GroupArg
	util.Bind(request, &arg)

	// 调用service 进行 添加好友 逻辑
	community, err := contactService.CreateGroup(&arg)

	if err != nil {
		util.RespError(writer, "添加好友失败:"+err.Error())
	} else {
		util.RespSuccess(writer, *community, "添加好友成功")
	}
}

// 加载群聊
func LoadCommunity(writer http.ResponseWriter, request *http.Request) {
	// 通过反射工具将 前端 传过来的 参数 绑定至 对象上
	var arg *args_user.UserTokenArg = &args_user.UserTokenArg{}
	util.Bind(request, arg)

	vos, err := contactService.LoadCommunity(arg)

	if nil != err {
		util.RespError(writer, "加载群列表失败:"+err.Error())
	} else {
		util.RespList(writer, 0, vos, len(vos))
	}
}

// 加入群聊
func JoinCommunity(writer http.ResponseWriter, request *http.Request) {
	// 通过反射工具将 前端 传过来的 参数 绑定至 对象上
	var arg args_contact.ContactGroupArg
	util.Bind(request, &arg)

	// 调用service 进行 添加群 逻辑
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)

	if err != nil {
		util.RespError(writer, "添加群聊失败:"+err.Error())
	} else {
		util.RespSuccess(writer, nil, "添加群聊成功")
	}
}
