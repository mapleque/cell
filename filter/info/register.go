package info

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/service/user"
)

// @author yangyang
// @review
// 通过手机号注册
// @filter
func Register(context *Context) bool {
	params := Map(context.Params["data"])
	mobile := String(params["username"])
	password := String(params["password"])
	userType := String(params["type"])

	_, err := user.RegisterByMobile(mobile, password, userType)
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
