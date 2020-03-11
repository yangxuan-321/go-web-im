package dao

import (
	"../model"
	"../util"
)

type UserDao struct {
}

// 通过用户号码查询 用户信息
func (*UserDao) FindUserByMobile(mobile string) (model.User, error) {
	user := model.User{}
	_, err := DbEngine.Where("mobile=?", mobile).Get(&user)
	return user, err
}

// 刷新token
func (*UserDao) RefreshToken(userInfo *model.User) (int64, error) {
	// 更新token
	token := util.GetUUID()
	(*userInfo).Token = token
	return DbEngine.ID((*userInfo).Id).Cols("token").Update(userInfo)
}

func (dao *UserDao) InsertOne(user *model.User) (int64, error) {
	return DbEngine.InsertOne(user)
}

func (dao *UserDao) FindUserByIds(friendIds []int64) ([]model.User, error) {
	users := make([]model.User, 0)
	err := DbEngine.In("id", friendIds).Find(&users)

	return users, err
}
