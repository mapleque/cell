package passport

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/service/system"
)

// @author yangyang
// @review
// 通过用户名密码获取token
// @filter
func Login(context *Context) bool {
	params := Map(context.Params["data"])
	username := String(params["username"])
	password := String(params["password"])

	// check username and password
	userId, err := system.CheckLogin(username, password)
	if err != 0 { // check login faild
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "Your username is: " + username
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "Your username is: " + username
			return RTBool(false)
		}
	}

	// userId to token
	token, _ := system.GeneralToken(userId)
	context.Data = token
	return true
}
