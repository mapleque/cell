package user

// @author yangyang
// @review
// 检查用户名密码
// @param username <用户名>
// @param password <密码>
// @return <userId>
// 		 > 0 success
// 		<= 0 faild
func CheckLogin(username, password string) int {
	if username == "yangyang" {
		return 1
	}
	return 0
}

// @author yangyang
// @review
// 生成token
// @param userId
// @return <token>
func GeneralToken(userId int) string {
	return "token"
}
