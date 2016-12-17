package student

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/tellus/constant"
)

// @author yangyang
// @review
// 获取用户信息
// @param userId <用户id>
// @param fields <要获取的字段>
// @return <info, err>
// 		err  =   0 success
// 				-1 db error
// 				 1 invalid user
//		info {} # with fields wanted
func GetInfo(
	userId int,
	fields []string,
	userType string) (map[string]interface{}, int) {
	conn := db.UseDB(DEF_DB)
	switch userType {
	case USER_TYPE_STUDENT:
		fieldMap := map[string][2]interface{}{
			"userId":    [2]interface{}{"user_id", 0},
			"studentId": [2]interface{}{"id", 0},
			"name":      [2]interface{}{"name", ""},
			"mobile":    [2]interface{}{"mobile", ""},
			"sex":       [2]interface{}{"sex", 0},
			"headImg":   [2]interface{}{"head_img", ""}}
		selectField := ""
		for _, field := range fields {
			if len(selectField) > 0 {
				selectField = selectField + ", "
			}
			selectField = selectField + String(fieldMap[field][0]) + " AS " + field
		}
		infos := conn.Select(
			`SELECT `+selectField+` FROM student WHERE user_id = ? LIMIT 1`,
			userId)
		if len(infos) != 1 {
			return nil, RTInt(-1)
		}
		for key, value := range infos[0] {
			if value == "null" {
				infos[0][key] = fieldMap[key][1]
			}
		}
		return infos[0], 0
	default:
		return nil, RTInt(1)
	}
}
