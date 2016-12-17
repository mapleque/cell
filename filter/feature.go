package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
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
					"feature_ids": []string{Rule(
						"int",
						STATUS_INVALID_ID,
						"要查的特权id")}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": []Checker{
					Checker{
						"key":   Rule("int", 0, "特权id"),
						"valid": Rule("int", 0, "特权id对应的状态"),
						"start": Rule("datetime", 0, "开始时间"),
						"end":   Rule("datetime", 0, "结束时间")}},
				"errmsg": "string"}},
			DefaultFilter)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新用户拥有的指定特权数量",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token"),
					"feature_id": Rule("int", STATUS_INVALID_ID, "特权id"),
					"start": Rule("datetime",
						STATUS_INVALID_TIME,
						"特权开始时间，空字符串不修改"),
					"end": Rule("datetime",
						STATUS_INVALID_TIME,
						"特权结束时间，空字符串不修改")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_MOBILE,
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"valid": Rule("int", 0, "修改后特权id对应的状态"),
					"start": Rule("datetime", 0, "修改后的开始时间"),
					"end":   Rule("datetime", 0, "修改后的结束时间")},
				"errmsg": "string"}},
			DefaultFilter)

		rt.NewDocRouter(&Doc{
			Path:        "manage",
			Description: "管理系统特权：添加"},
			DefaultFilter)
	})
}
