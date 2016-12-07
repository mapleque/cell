package filter

import (
	. "github.com/coral"

	. "github.com/tellus/constant"
	"github.com/tellus/filter/passport"
)

// @author yangyang
// @review
// 初始化passport的router
func init() {
	RegisterFilter("passport", func(rt *Router) {
		rt.NewDocRouter(&Doc{
			Path:        "login",
			Description: "用户登录",
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
			passport.Login)

		rt.NewDocRouter(&Doc{
			Path:        "logout",
			Description: "用户登出",
			Input: Checker{
				"data": Checker{
					"token": Rule(
						"string",
						STATUS_INVALID_TOKEN,
						"用户token")}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER),
				"data":   "string",
				"errmsg": "string"}},
			DefaultFilter)

	})
}
