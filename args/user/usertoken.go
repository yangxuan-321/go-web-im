package args_user

type UserTokenArg struct {
	Userid int64  `json:"userid" form:"userid"`
	Token  string `json:"token" form:"dstid"`
}
