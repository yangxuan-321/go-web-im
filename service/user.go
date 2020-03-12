package service

import (
	"../dao"
	"../model"
	"../util"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type UserService struct {
}

var userDao *dao.UserDao = &dao.UserDao{}

/**
 * mobile - 手机
 * plainpwd - 明文密码
 * nickname - 昵称
 * avastar
 */
func (*UserService) Register(mobile, plainpwd, niclname, avatar, sex string) (model.User, error) {
	// 检测手机号 是否存在
	// 传入 用户 信息的 实例化 地址, 用于 存放返回数据
	userInfo, err_ := userDao.FindUserByMobile(mobile)
	if nil != err_ {
		return userInfo, err_
	}

	// 如果存在 则返回提示 已经注册
	if userInfo.Id > 0 {
		return userInfo, errors.New("该手机号已经被注册过")
	}

	// 否则 拼接 插入数据库
	userInfo.Mobile = mobile
	userInfo.Avatar = avatar
	userInfo.Nickname = niclname
	userInfo.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	userInfo.Passwd = util.MakePasswd(plainpwd, userInfo.Salt)
	userInfo.Sex = sex
	userInfo.Createat = time.Now()
	//token 可以是一个随机数
	userInfo.Token = util.GetUUID()

	total, err := userDao.InsertOne(&userInfo)
	if nil != err || 0 == total {
		log.Fatal("插入不成功")
	}

	// 返回新用户信息
	return userInfo, nil
}

func (*UserService) Login(mobile, plainpwd string) (model.User, error) {
	// 通过手机号查询用户
	userInfo, err := userDao.FindUserByMobile(mobile)
	if userInfo.Id == 0 {
		return userInfo, errors.New("用户不存在")
	}

	// 查询到了比对密码
	if !util.ValidatePasswd(plainpwd, userInfo.Salt, userInfo.Passwd) {
		return userInfo, errors.New("密码不正确")
	}

	// 更新token
	token := util.GetUUID()
	userInfo.Token = token
	updateCount, err := userDao.RefreshToken(&userInfo)
	if err != nil || updateCount < 1 {
		return userInfo, errors.New("更新token失败:" + err.Error())
	}

	return userInfo, nil
}

// 通过 Id 查询 User
func (service *UserService) FindUserById(userId int64) (model.User, error) {
	return userDao.FindUserById(userId)
}
