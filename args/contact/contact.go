package args_contact

import (
	"../../args"
)

type ContactArg struct {
	args.PageArg
	Userid    int64  `json:"userid" form:"userid"`
	Dstmobile string `json:"dstmobile" form:"dstmobile"`
}
