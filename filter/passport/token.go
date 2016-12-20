package passport

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	"github.com/cell/service/user"
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
			context.Status = STATUS_INVALID_PASSWORD
			context.Errmsg = "Your username is: " + username
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}
	ret := map[string]interface{}{"token": token}
	context.Data = ret
	return true
}

// @author yangyang
// @review
// 注销用户登录token
// @filter
func Logout(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	userId, checkTokenCode := user.CheckToken(token)
	if checkTokenCode != 0 {
		switch checkTokenCode {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_INVALID_TOKEN
			context.Errmsg = "Your token is: " + token
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}
	_, logoutCode := user.Logout(userId)
	if logoutCode != 0 {
		switch logoutCode {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_INVALID_TOKEN
			context.Errmsg = "Your token is: " + token
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
