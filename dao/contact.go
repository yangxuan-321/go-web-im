package dao

import (
	"../model"
	"time"
)

type ContactUserDao struct {
}

func (dao *ContactUserDao) FindFriendsByOwneridAndDestid(userId int64, destId int64) (model.ContactUser, error) {
	contactUserInfo := model.ContactUser{}
	_, err := DbEngine.Where("ownerid = ?", userId).And("destid=?", destId).Get(&contactUserInfo)
	return contactUserInfo, err
}

func (dao *ContactUserDao) Addfriends(userid, destid int64) error {
	// 启动事物
	session := DbEngine.NewSession()
	session.Begin()

	// 添加两条好有记录
	//userid -> destid 此处的备注 需要优化 默认为 用户的用户名
	contactUser1 := model.ContactUser{Ownerid: userid, Dstobj: destid, Memo: "", Createat: time.Now()}
	//destid -> userid
	contactUser2 := model.ContactUser{Ownerid: destid, Dstobj: userid, Memo: "", Createat: time.Now()}

	inserCount, err := session.Insert(contactUser1, contactUser2)
	if err != nil && inserCount < 2 {
		session.Rollback()
		return err
	}

	session.Commit()

	return nil
}

func (dao *ContactUserDao) FindFriendsByOwnerid(userid interface{}) ([]model.ContactUser, error) {
	contactUsers := make([]model.ContactUser, 0)
	err := DbEngine.Where("ownerid = ?", userid).Find(&contactUsers)
	return contactUsers, err
}
