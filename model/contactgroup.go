package model

import "time"

// 群聊放在这个表
type ContactGroup struct {
	Id         int64     	`xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000 自己的id
	Ownerid       int64		`xorm:"bigint(20)" form:"ownerid" json:"ownerid"`   // 记录是谁的
	//对端,10001 好友的id
	Groupid       int64		`xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对端信息
	//
	Cate      int			`xorm:"int(11)" form:"cate" json:"cate"`   // 什么类型
	Memo    string			`xorm:"varchar(120)" form:"memo" json:"memo"`   // 备注
	//
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 创建时间
}
