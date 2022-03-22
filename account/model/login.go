package model

type LoginReq struct {
	PassportUid  int64  `form:"passport_uid" json:"passport_uid"`
	PassportCode string `form:"passport_code" json:"passport_code"`
}

const (
	LoginResultCodeCorrect = 0 // 登录成功
	LoginResultCodeFailed = 1 // 系统内部错误
	LoginReusulCodeError = 2 // 密码错误
)
