package user

import (
	"strings"

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
	//TODO change to config get key
	// get user password
	return common.AesEcbEnc("yestem11yestem11yestem11yestem11", String(userId)+"||password"), 0
}

// @author yangyang
// @review
// 检查token
// @param token
// @return <mobile, err>
// 		err  =   0 success
// 				 1 invalid user
func CheckToken(token string) (string, int) {
	//TODO change to config get key
	studentStr := String(common.AesEcbDec("yestem11yestem11yestem11yestem11", token))
	if len(studentStr) < 1 {
		return "", 1
	}
	studentArr := strings.Split(studentStr, "||")
	if len(studentArr) < 2 {
		return "", 1
	}
	// TODO check password
	mobile := studentArr[0]

	return mobile, 0
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
func GetInfo(mobile string, field []string) (map[string]interface{}, int) {
	if mobile != "" {
		ret := make(map[string]interface{})
		ret["studentId"] = 1
		ret["userId"] = 1
		ret["mobile"] = mobile
		return ret, 0
	}
	return nil, 1
}
