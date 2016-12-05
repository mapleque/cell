package passport

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	user "github.com/tellus/service/user"
)

// @author yangyang
// @review
// 通过用户名密码获取token
// @filter
func GetToken(context *Context) bool {
	params := Map(context.Params["data"])
	username := String(params["username"])
	password := String(params["password"])

	// check username and password
	userId, err := user.CheckLogin(username, password)
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
	token, _ := user.GeneralToken(userId)
	context.Data = token
	return true
}

// @author yangyang
// @review
// 通过token获取用户信息
// @filter
func GetInfo(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	fields := Array(params["fields"])
	var wantFields []string
	for _, f := range fields {
		wantFields = append(wantFields, String(f))
	}

	// token to userId
	mobile, err := user.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}

	// get user info
	info, err := user.GetInfo(mobile, wantFields)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "Your userId is: " + mobile
			return RTBool(false)
		}
	}
	context.Data = info
	return true
}
