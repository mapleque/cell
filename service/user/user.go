package user

import (
	. "github.com/coral"
	"github.com/tellus/common"
)

// @author yangyang
// @review
// 检查用户名密码
// @param username <用户名>
// @param password <密码>
// @return <userId, err>
// 		err  =   0 success
// 				-1 db error
// 				 1 invalid user
func CheckLogin(username, password string) (int, int) {
	if username == "13800138000" {
		return 1, 0
	}
	return 0, 1
}

// @author yangyang
// @review
// 生成token
// @param userId
// @return <token, err>
// 		err  =   0 success
func GeneralToken(userId int) (string, int) {
	return common.AesEcbEnc("yestem11yestem11yestem11yestem11", String(userId)), 0
}

// @author yangyang
// @review
// 检查token
// @param token
// @return <userId, err>
// 		err  =   0 success
// 				 1 invalid user
func CheckToken(token string) (int, int) {
	userId := Int(common.AesEcbDec("yestem11yestem11yestem11yestem11", token))

	if userId > 0 {
		return userId, 0
	}
	return 0, 1
}

// @author yangyang
// @review
// 获取用户信息
// @param userId <用户id>
// @param field <要获取的字段>
// @return <info, err>
// 		err  =   0 success
// 				-1 db error
// 				 1 invalid user
//		info {
//			userId:
//			studentId:
//		}
func GetInfo(userId int, field []string) (map[string]interface{}, int) {
	if userId == 1 {
		ret := make(map[string]interface{})
		ret["studentId"] = 1
		ret["userId"] = 1
		return ret, 0
	}
	return nil, 1
}
