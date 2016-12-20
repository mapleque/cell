package resource

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	service "github.com/cell/service/resource"
	"github.com/cell/service/user"
)

// @author yangyang
// @review
// 查询用户当前系统资源
// @filter
func Check(context *Context) bool {
	params := Map(context.Params["data"])
	token := String(params["token"])
	productIds := Array(params["product_ids"])
	var ids []int
	for _, productId := range productIds {
		ids = append(ids, Int(productId))
	}

	// token to userId
	userId, err := user.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}
	info, err := service.GetAmountBatch(userId, ids)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}

	context.Data = info
	return true
}
