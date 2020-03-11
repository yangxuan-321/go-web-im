package vo

type UserFriendsVO struct {
	// 登录用户id
	Userid int64 `json:"id"`

	// 好友/群 id
	Destid int64 `json:"id"`

	// 好友类型 - 群/单人
	Friendtype string `json:"friend_type"`

	// 备注
	Memo string `json:"memo"`

	// 头像
	Avatar string `json:"avatar"`
}
