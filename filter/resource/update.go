package resource

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	service "github.com/cell/service/resource"
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
	amount := Int(params["amount"])

	// token to userId
	userId, err := user.CheckToken(token)
	if err != 0 { // check token faild
		context.Status = STATUS_ERROR_INVALID_USER
		context.Errmsg = "Your token is: " + token
		return RTBool(false)
	}
	tarAmount, err := service.UpdateAmount(userId, productId, amount)
	if err != 0 {
		switch err {
		case -1:
			context.Status = STATUS_ERROR_DB
			return RTBool(false)
		case 1:
			context.Status = STATUS_INVALID_AMOUNT
			context.Errmsg = "Your target amount is: " + String(tarAmount)
			return RTBool(false)
		case 2:
			context.Status = STATUS_INVALID_ID
			context.Errmsg = "Your target product id is: " + String(productId)
			return RTBool(false)
		default:
			context.Status = STATUS_ERROR_UNKNOWN
			context.Errmsg = "unexpect error"
			return RTBool(false)
		}
	}

	ret := map[string]interface{}{"amount": tarAmount}
	context.Data = ret
	return true
}
