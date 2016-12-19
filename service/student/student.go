package student

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/tellus/constant"
)

// @autho yangyang
// @review
// 用于本模块的常量定义
var FIELDS_MAP = map[string][2]interface{}{
	"user_id":    [2]interface{}{"user_id", 0},
	"student_id": [2]interface{}{"id", 0},
	"name":       [2]interface{}{"name", ""},
	"mobile":     [2]interface{}{"mobile", ""},
	"sex":        [2]interface{}{"sex", 0},
	"head_img":   [2]interface{}{"head_img", ""}}

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
	fields []string) (map[string]interface{}, int) {
	selectField := ""
	for _, field := range fields {
		if len(selectField) > 0 {
			selectField = selectField + ", "
		}
		selectField = selectField +
			String(FIELDS_MAP[field][0]) + " AS " + field
	}
	conn := db.UseDB(DEF_DB)
	infos := conn.Select(
		`SELECT `+selectField+` FROM student WHERE user_id = ? LIMIT 1`,
		userId)
	if len(infos) != 1 {
		return nil, RTInt(-1)
	}
	for key, value := range infos[0] {
		if value == "null" {
			infos[0][key] = FIELDS_MAP[key][1]
		}
	}
	return infos[0], 0
}

// @author yangyang
// @review
// 修改用户信息
// @param userId <用户id>
// @param fields <要修改的字段>
// @return err
// 		err  =   0 success
// 				-1 db error
// 				 1 invalid user
func UpdateInfo(userId int, fields map[string]interface{}) int {
	updateField := ""
	var paramField []interface{}
	for key, value := range fields {
		if len(updateField) > 0 {
			updateField = updateField + ","
		}
		updateField = updateField + String(FIELDS_MAP[key][0]) + " = ?"
		paramField = append(paramField, value)
	}
	paramField = append(paramField, userId)
	conn := db.UseDB(DEF_DB)
	ret := conn.Update(
		`UPDATE student SET `+updateField+` WHERE user_id = ? LIMIT 1`,
		paramField...)
	if ret != 1 {
		return RTInt(-1)
	}
	return 0
}
