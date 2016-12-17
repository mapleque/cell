package product

import (
	. "github.com/coral"

	service "github.com/tellus/service/product"
)

// @author yangyang
// @review
// 添加一个新的系统资源
// @filter
func New(context *Context) bool {
	params := Map(context.Params["data"])
	productType := String(params["type"])
	name := String(params["name"])
	additional := String(params["additional"])

	id, err := service.New(productType, name, additional)
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

	context.Data = map[string]interface{}{"id": id}
	return true
}
