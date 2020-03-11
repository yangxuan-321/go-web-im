package model

import "time"

//好友存在这个表里面
type ContactUser struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000 自己的id
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//对端,10001 好友的id
	Dstobj int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	//备注
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"`
	//添加时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
