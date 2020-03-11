package model

import "time"

// 群聊放在这个表
type ContactGroup struct {
	//主键
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//自己的ID
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//加入的群的ID
	Groupid int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	//群备注
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"`
	//入群时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
