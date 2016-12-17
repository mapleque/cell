package product

import (
	. "github.com/coral"

	service "github.com/tellus/service/product"
)

// @author yangyang
// @review
// 查询系统资源列表
// @filter
func List(context *Context) bool {
	params := Map(context.Params["data"])
	pageSize := Int(params["pageSize"])
	pageStart := Int(params["pageStart"])

	info, err := service.List(pageSize, pageStart)
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
