package resource

import (
	. "github.com/coral"
	"github.com/coral/db"

	. "github.com/tellus/constant"
)

// @autho yangyang
// @review
// 用于本模块的常量定义
var FIELDS_MAP = map[string][2]interface{}{
	"userId":    [2]interface{}{"user_id", 0},
	"studentId": [2]interface{}{"id", 0},
	"name":      [2]interface{}{"name", ""},
	"mobile":    [2]interface{}{"mobile", ""},
	"sex":       [2]interface{}{"sex", 0},
	"headImg":   [2]interface{}{"head_img", ""}}
