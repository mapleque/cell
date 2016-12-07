package info

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/service/info"
	"github.com/tellus/service/system"
)

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
	mobile, err := system.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}

	// get user info
	info, err := info.GetInfo(mobile, wantFields)
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
