package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/filter/feature"
)

// @author yangyang
// @review
// 初始化feature的router
func init() {
	RegisterFilter("feature", func(rt *Router) {
		rt.NewDocRouter(&Doc{
			Path:        "check",
			Description: "检查用户指定特权状态",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token"),
					"product_ids": []string{Rule(
						"int",
						STATUS_INVALID_ID,
						"要查的特权id")}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": []Checker{
					Checker{
						"product_id": Rule("int", 0, "特权id"),
						"valid":      Rule("int", 0, "特权id对应的状态"),
						"start":      Rule("datetime", 0, "开始时间"),
						"end":        Rule("datetime", 0, "结束时间")}},
				"errmsg": "string"}},
			feature.Check)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新用户拥有的指定特权数量",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token"),
					"product_id": Rule("int", STATUS_INVALID_ID, "特权id"),
					"start": Optional(Rule("datetime",
						STATUS_INVALID_TIME,
						"特权开始时间，空字符串不修改")),
					"end": Optional(Rule("datetime",
						STATUS_INVALID_TIME,
						"特权结束时间，空字符串不修改"))}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_MOBILE,
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"start": Optional(Rule("datetime", 0, "修改后的开始时间")),
					"end":   Optional(Rule("datetime", 0, "修改后的结束时间"))},
				"errmsg": "string"}},
			feature.Update)
	})
}
