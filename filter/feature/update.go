package feature

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	service "github.com/cell/service/feature"
	"github.com/cell/service/user"
)

// @author yangyang
// @review
// 更新用户当前系统资源
// @filter
func Update(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	productId := Int(params["product_id"])
	start := ""
	end := ""
	if params["start"] != nil {
		start = String(params["start"])
	}
	if params["end"] != nil {
		end = String(params["end"])
	}
	if start == "" && end == "" {
		context.Status = STATUS_INVALID_TIME
		return RTBool(false)
	}

	// token to userId
	userId, err := user.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}
	err = service.Update(userId, productId, start, end)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_INVALID_ID
			context.Errmsg = "Your target product id is: " + String(productId)
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}

	ret := make(map[string]interface{})
	if start != "" {
		ret["start"] = start
	}
	if end != "" {
		ret["end"] = end
	}
	context.Data = ret
	return true
}
