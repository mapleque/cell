package info

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/service/student"
	"github.com/tellus/service/user"
)

// @author yangyang
// @review
// 通过token获取用户信息
// @filter
func GetStudentInfo(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	fields := Array(params["fields"])
	if len(fields) < 1 {
		context.Status = STATUS_INVALID_FIELD
		context.Errmsg = "Your field is none"
		return RTBool(false)
	}
	var wantFields []string
	for _, f := range fields {
		wantFields = append(wantFields, String(f))
	}

	// token to userId
	userId, err := user.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}

	// get user info
	info, err := student.GetInfo(userId, wantFields)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_ERROR_INVALID_USER
			context.Errmsg = "Your userId is: " + String(userId)
			return RTBool(false)
		}
	}
	context.Data = info
	return true
}
