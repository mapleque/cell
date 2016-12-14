package passport

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/service/user"
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
	token, code := user.Login(username, password)
	if code != 0 {
		switch code {
		case 1:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "Your username is: " + username
			return RTBool(false)
		case 2:
			context.Status = STATUS_ERROR_INVALID_PASSWORD
			context.Errmsg = "Your username is: " + username
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}
	context.Data = token
	return true
}

// @author yangyang
// @review
// 注销用户登录token
// @filter
func Logout(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	userId, code := user.CheckToken(token)
	if code != 0 {
		switch code {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_TOKEN
			context.Errmsg = "Your token is: " + token
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}
	_, code := user.Logout(userId)
	if code != 0 {
		switch code {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_TOKEN
			context.Errmsg = "Your token is: " + token
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}
	Context.data = ""
	return true
}
