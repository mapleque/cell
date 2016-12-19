package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/filter/resource"
)

// @author yangyang
// @review
// 初始化resource的router
func init() {
	RegisterFilter("resource", func(rt *Router) {
		rt.NewDocRouter(&Doc{
			Path:        "check",
			Description: "检查用户指定资源数量",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token"),
					"product_ids": []string{Rule(
						"int",
						STATUS_INVALID_ID,
						"要查的资源id")}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_ID,
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": []Checker{
					Checker{
						"product_id": Rule("int", 0, "资源id"),
						"amount":     Rule("int", 0, "资源id对应的数量")}},
				"errmsg": "string"}},
			resource.Check)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新用户拥有的指定资源数量",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token"),
					"product_id": Rule(
						"int",
						STATUS_INVALID_ID,
						"资源id"),
					"amount": Rule(
						"int",
						STATUS_INVALID_AMOUNT,
						"资源变化数量")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_INVALID_ID,
					STATUS_INVALID_AMOUNT,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"amount": Rule("int", 0, "资源变化之后的数量")},
				"errmsg": "string"}},
			resource.Update)
	})
}
