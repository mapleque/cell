package server

import "errors"

var (
	errNoAuthToken          = errors.New("无法找到登录凭证")
	errAuthTokenInvalid     = errors.New("登录凭证已失效")
	errCaptchaInvalid       = errors.New("验证码不正确")
	errUserExist            = errors.New("该用户已注册")
	errUserNotExist         = errors.New("该用户未注册")
	errTGTInvalid           = errors.New("错误的tgt格式")
	errSTInvalid            = errors.New("错误的st格式")
	errAuthenticatorInvalid = errors.New("错误的authenticator")
	errAppNotExist          = errors.New("应用未注册")
	errEncDataInvalid       = errors.New("错误的密文数据")
	errSendMail             = errors.New("邮件发送错误，请稍后重试")
	errAPPNameInvalid       = errors.New("应用名称不能为空")
)
