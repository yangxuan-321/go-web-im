package model

import "time"

const (
	// 女
	SEX_WOMEN="W"
	// 男
	SEX_MEN="M"
	// 未知
	SEX_UNKNOW="U"
)

type User struct {
	//用户ID
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	// 手机
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	// 密码
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`
	// 头像
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	// 性别
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	// 昵称
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// 加盐随机字符串6
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`
	// 是否在线
	Online     int	`xorm:"int(10)" form:"online" json:"online"`
	// 前端鉴权因子,
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`
	// 备注
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`
	// 创建时间
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`
}