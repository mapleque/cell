package info

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	"github.com/cell/service/user"
)

// @author yangyang
// @review
// 通过手机号注册
// @filter
func Register(context *Context) bool {
	params := Map(context.Params["data"])
	mobile := String(params["username"])
	password := String(params["password"])

	_, err := user.RegisterByMobile(mobile, password)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "User mobile is exist: " + mobile
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}

	context.Data = ""
	return true
}
