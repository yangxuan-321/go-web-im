package args_contact

import (
	"../../args"
)

type ContactUserArg struct {
	args.PageArg
	Userid    int64  `json:"userid" form:"userid"`
	Dstmobile string `json:"dstmobile" form:"dstmobile"`
}

type ContactGroupArg struct {
	args.PageArg
	Userid int64  `json:"userid" form:"ownerid"`
	Dstid  string `json:"dstid" form:"dstid"`
}
