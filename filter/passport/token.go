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
	params := context.Params
	username := params["username"].(string)
	password := params["password"].(string)
	userId := user.CheckLogin(username, password)
	if userId > 0 {
		token := user.GeneralToken(userId)
		context.Data = token
		return true
	}
	context.Status = STATUS_INVALID_USER
	context.Errmsg = username
	return false
}

func GetInfo(context *Context) bool {
	return true
}
