package args_contact

const (
	GroupType_默认   = 0x00
	GroupType_兴趣爱好 = 0x01
	GroupType_行业交流 = 0x02
	GroupType_生活休闲 = 0x03
	GroupType_学习考试 = 0x04
)

type GroupArg struct {
	// 头像
	Avatar string `json:"avatar" form:"icon"`
	// 群组类型 - 参照上面枚举
	Grouptype int `json:"grouptype" form:"cate"`
	// 群组名称
	Groupname string `json:"groupname" form:"name"`
	// 群组介绍
	Groupdesc string `json:"groupdesc" form:"memo"`
	// 创建人
	Ownerid int64 `json:"ownerid" form:"ownerid"`
}
