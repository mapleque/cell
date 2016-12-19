package user

import (
	"strings"

	. "github.com/coral"
	"github.com/coral/config"
	"github.com/coral/db"

	"github.com/tellus/common"
	. "github.com/tellus/constant"
)

// @author yangyang
// @review
// 检查用户名密码
// @param mobile <手机号>
// @param password <密码>
// @return <token, code>
// 		err  =   0 success
// 				 1 invalid user
// 				 2 invalid password
func Login(mobile, password string) (string, int) {
	conn := db.UseDB(DEF_DB)
	// 先查手机号是否存在
	users := conn.Select(
		`SELECT id, password FROM user WHERE username = ? LIMIT 1`,
		mobile)
	if len(users) != 1 {
		return "", RTInt(1)
	}
	encPassword := String(users[0]["password"])
	// 再检查密码是否正确
	if !checkPassword(password, encPassword) {
		return "", RTInt(2)
	}
	conf := config.Use(DEF_CONF)
	key := conf.String("token.USER_TOKEN_KEY")
	token := common.AesEcbEnc(key, mobile+"||"+encPassword)
	// TODO rand USER_TOKEN_KEY 并缓存
	return token, 0
}

// @author yangyang
// @review
// 用户退出登录
// @param userId
// @return <token, code>
// 		err  =   0 success
// 				 1 invalid user
func Logout(userId int) (string, int) {
	// TODO 清除 USER_TOKEN_KEY 缓存
	return "", 0
}

// @author yangyang
// @review
// 检查token
// @param token
// @return <userId, code>
// 		err  =   0 success
// 				 1 invalid token
func CheckToken(token string) (int, int) {
	conf := config.Use(DEF_CONF)
	// TODO 从缓存取 USER_TOKEN_KEY
	key := conf.String("token.USER_TOKEN_KEY")
	studentStr := String(common.AesEcbDec(key, token))
	if len(studentStr) < 1 {
		return 0, RTInt(1)
	}
	studentArr := strings.Split(studentStr, "||")
	if len(studentArr) < 2 {
		return 0, RTInt(1)
	}
	mobile := studentArr[0]
	password := studentArr[1]

	// TODO 延长 USER_TOKEN_KEY 缓存过期时间
	conn := db.UseDB(DEF_DB)
	users := conn.Select(
		`SELECT id, password FROM user WHERE username = ? LIMIT 1`,
		mobile)
	if len(users) != 1 {
		return 0, RTInt(1)
	}

	if password != String(users[0]["password"]) {
		return 0, RTInt(1)
	}

	return Int(users[0]["id"]), 0
}

// @author yangyang
// @review
// 通过手机号注册
// @param username <手机号>
// @param password <密码>
// @param userType <用户类型, USER_TYPE_*>
// @return <userId, code>
// 		err  =   0 success
// 				-1 db error
// 				 1 invalid user
func RegisterByMobile(mobile, password, userType string) (int, int) {
	conn := db.UseDB(DEF_DB)
	// 先查手机号是否存在
	users := conn.Select(
		`SELECT id FROM user WHERE username = ? LIMIT 1`,
		mobile)
	if len(users) > 0 {
		return 0, RTInt(1)
	}

	// 密码加密
	token := encryptPassword(password)

	trans := conn.Begin()
	// 插入system，记录userId
	userId := conn.Insert(
		`INSERT INTO user (username, password, status) VALUES (?,?,?)`,
		mobile, token, USER_STATUS_NORMAL)
	if userId <= 0 {
		trans.Rollback()
		return 0, RTInt(-1)
	}

	// 插入对应类型user info
	switch userType {
	case USER_TYPE_STUDENT:
		studentId := conn.Insert(
			`INSERT INTO student (user_id, mobile) VALUES (?,?)`,
			userId, mobile)
		if studentId <= 0 {
			trans.Rollback()
			return 0, RTInt(-1)
		}
	default:
		return 0, RTInt(2)
	}

	trans.Commit()
	// 返回userId
	return Int(userId), 0
}

// @author yangyang
// @review
// 加密用户密码
// 加密方式：
// 1.随机生成salt
// 2.对用户密码加入salt后再md5产生token
// 3.将token和salt同时入库存储
// @param password <用户密码>
// @return token <入库的token，格式为encPassword|salt>
func encryptPassword(password string) string {
	salt := common.RandMD5()
	encPassword := common.MD5(password + "__" + salt)
	return encPassword + "|" + salt
}

// @author yangyang
// @review
// 校验用户密码
// @param password <用户密码>
// @param token <存储的token>
// @return <校验成功或失败>
func checkPassword(password, token string) bool {
	tmpArr := strings.Split(token, "|")
	if len(tmpArr) > 1 {
		encPassword := tmpArr[0]
		salt := tmpArr[1]
		return common.MD5(password+"__"+salt) == encPassword
	} else {
		return password == token
	}
}
