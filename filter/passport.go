package filter

import (
	coral "github.com/coral"

	passport "github.com/tellus/filter/passport"
)

// @author yangyang
// @review
// 初始化passport的router
func init() {
	register("passport", func(rt *coral.Router) {
		rt.NewDocRouter(&coral.Doc{
			Path:        "token",
			Description: "获取用户token",
			Input: coral.Checker{
				"data": coral.Checker{
					"username": "mobile",
					"password": "md5"}},
			Output: coral.Checker{
				"status": "int",
				"data": coral.Checker{
					"token": "md5"},
				"errmsg": "string"}},

			passport.GetToken)

		rt.NewDocRouter(&coral.Doc{
			Path:        "info",
			Description: "获取用户信息",
			Input: coral.Checker{
				"data": coral.Checker{
					"token":  "md5",
					"fields": []string{"string{userId,studentId}"}}},
			Output: coral.Checker{
				"status": "int",
				"data": coral.Checker{
					"token": "md5"},
				"errmsg": "string"}},

			passport.GetInfo)
	})
}
