package model

import "time"

type Community struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//群名称
	Name string `xorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//群logo
	Icon string `xorm:"varchar(250)" form:"icon" json:"icon"`
	//群类型
	Type int `xorm:"bigint(20)" form:"type" json:"type"`
	//描述
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"`
	//群创建时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
