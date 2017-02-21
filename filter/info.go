package filter

import (
	. "github.com/coral"

	. "github.com/cell/constant"
	"github.com/cell/filter/info"
)

// @author yangyang
// @review
// 初始化info的router
// info/register
// info/check
// info/update
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
				"data":   "string",
				"errmsg": "string"}},
			info.Register)

		rt.NewDocRouter(&Doc{
			Path:        "check",
			Description: "获取用户信息，只返回fieds中指定的信息字段",
			Input: Checker{
				"data": Checker{
					"token": Rule("string", STATUS_INVALID_TOKEN, "用户token"),
					"fields": []string{Rule(
						"string{user_id,role_id,name,mobile,sex,head_img}",
						STATUS_INVALID_FIELD,
						"指定要获取的字段")}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_INVALID_FIELD,
					STATUS_ERROR_INVALID_USER),
				"data": Checker{
					"user_id":  Optional(Rule("int", 0, "用户id")),
					"role_id":  Optional(Rule("int", 0, "角色_id")),
					"name":     Optional(Rule("string", 0, "姓名")),
					"mobile":   Optional(Rule("string", 0, "手机号")),
					"sex":      Optional(Rule("int", 0, "性别")),
					"head_img": Optional(Rule("string", 0, "头像url"))},
				"errmsg": "string"}},

			info.GetInfo)

		rt.NewDocRouter(&Doc{
			Path:        "update",
			Description: "更新用户信息",
			Input: Checker{
				"data": Checker{
					"token": Rule("string", STATUS_INVALID_TOKEN, "用户token"),
					"fields": Checker{
						"name": Optional(
							Rule("string", STATUS_INVALID_NAME, "用户姓名")),
						"mobile": Optional(
							Rule("string", STATUS_INVALID_MOBILE, "用户姓名")),
						"head_img": Optional(
							Rule("string", STATUS_INVALID_URL, "用户姓名")),
						"sex": Optional(
							Rule(InInt(
								USER_SEX_MALE,
								USER_SEX_FEMALE),
								STATUS_INVALID_SEX, "用户性别"))}}},
			Output: Checker{
				"status": InStatus(
					STATUS_INVALID_TOKEN,
					STATUS_ERROR_INVALID_USER,
					STATUS_INVALID_FIELD,
					STATUS_INVALID_NAME,
					STATUS_INVALID_MOBILE,
					STATUS_INVALID_URL,
					STATUS_INVALID_SEX),
				"data":   "string",
				"errmsg": "string"}},
			info.UpdateInfo)
	})
}
