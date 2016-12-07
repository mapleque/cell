package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/filter/info"
)

// @author yangyang
// @review
// 初始化info的router
func init() {
	RegisterFilter("info", func(rt *Router) {
		rt.NewDocRouter(&Doc{
			Path:        "register",
			Description: "用户注册",
			Input: Checker{
				"data": Checker{
					"username": Rule(
						"mobile",
						STATUS_INVALID_MOBILE,
						"用户手机号"),
					"password": Rule(
						"md5",
						STATUS_INVALID_TOKEN,
						"用户密码")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_MOBILE,
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"token": "string"},
				"errmsg": "string"}},
			DefaultFilter)

		rt.NewDocRouter(&Doc{
			Path:        "check",
			Description: "获取用户信息，只返回fieds中指定的信息字段",
			Input: Checker{
				"data": Checker{
					"token": Rule("string", STATUS_INVALID_TOKEN, "用户token"),
					"fields": []string{Rule(
						"string{userId,studentId}",
						STATUS_INVALID_FIELD,
						"指定要获取的字段")}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_INVALID_FIELD,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"userId":    Rule("int", 0, "如果fields未指定，值为0"),
					"studentId": Rule("int", 0, "如果fields未指定，值为0"),
					"mobile": Rule(
						"mobile",
						0,
						"如果fields未指定，值为空字符串")},
				"errmsg": "string"}},

			info.GetInfo)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新用户信息"},
			DefaultFilter)
	})
}
