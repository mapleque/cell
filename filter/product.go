package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/filter/product"
)

// @author yangyang
// @review
// 初始化product的router
func init() {
	RegisterFilter("product", func(rt *Router) {
		rt.NewDocRouter(&Doc{
			Path:        "list",
			Description: "获取系统资源列表",
			Input: Checker{
				"data": Checker{
					"pageSize": Rule(
						"int[1,-1]",
						STATUS_INVALID_PAGE,
						"单页数量"),
					"pageStart": Rule(
						"int[0,-1]",
						STATUS_INVALID_PAGE,
						"上次查询的最大id")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_PAGE),
				"data": []Checker{
					Checker{
						"id":          Rule("int", 0, "资源id"),
						"type":        Rule("string", 0, "资源类型"),
						"name":        Rule("string", 0, "资源名称"),
						"additional":  Rule("string", 0, "资源描述"),
						"update_time": Rule("string", 0, "资源最后更新时间"),
						"create_time": Rule("string", 0, "资源首次添加时间")}},
				"errmsg": "string"}},
			product.List)

		rt.NewDocRouter(&Doc{
			Path:        "new",
			Description: "新建系统资源",
			Input: Checker{
				"data": Checker{
					"type": Rule(
						InString(
							PRODUCT_TYPE_RESOURCE,
							PRODUCT_TYPE_FEATURE),
						STATUS_INVALID_TYPE,
						"资源类型"),
					"name": Rule(
						"string[1,20]",
						STATUS_INVALID_NAME,
						"资源名称"),
					"additional": Rule(
						"string[0,1000]",
						STATUS_INVALID_TEXT,
						"资源描述")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TYPE,
					STATUS_INVALID_NAME,
					STATUS_INVALID_TEXT),
				"data": Checker{
					"id": "int"},
				"errmsg": "string"}},
			product.New)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新系统资源",
			Input: Checker{
				"data": Checker{
					"id": Rule(
						"int[1,-1]",
						STATUS_INVALID_ID,
						"资源id"),
					"name": Rule(
						"string[1,20]",
						STATUS_INVALID_NAME,
						"资源名称"),
					"additional": Rule(
						"string[0,1000]",
						STATUS_INVALID_TEXT,
						"资源描述")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TYPE,
					STATUS_INVALID_NAME,
					STATUS_INVALID_TEXT),
				"data":   "string",
				"errmsg": "string"}},
			product.Update)

	})
}
