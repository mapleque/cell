package info

import (
//	. "github.com/coral"
)

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
