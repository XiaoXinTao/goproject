package consts
type ErrCode int

const (
	ErrSuccess        = 0
	ErrParams         = -1001
	ErrTokenInvalid   = -1002
	ErrTokenTimeout   = -1003
	ErrRecordNotFound = -1004
	ErrSignNotMatch   = -1011
	ErrTokenNotMatch  = -1012
	ErrConfig         = -2001
	ErrShark          = -2002
	ErrDB             = -2003

	ErrSystem = -5000
)

type ErrMsg string

const (
	ErrStrSuccess = "success"
	ErrStrFailure = "failure"

	ErrStrSystem                = "system error, please try again"
	ErrStrDB                    = "db error, please try again"
	ErrStrData                  = "data error"
	ErrStrInvalidParams         = "invalid or missing parameters"
	ErrStrInvalidParamPrefix    = "invalid parameter: "
	ErrStrSignNotMatch          = "sign not match"
	ErrStrTokenInvalid          = "token invalid"
	ErrStrTokenTimeout          = "token timeout"
	ErrStrUserNotFound          = "user not found"
	ErrStrChannelNotFound       = "channel not found"
	ErrStrPIPO                  = "pipo error"
	ErrStrGoodsNotFound         = "goods not found"
	ErrStrGameNotFound          = "game not found"
	ErrStrOrderNotFound         = "order not found"
	ErrStrRequestBlocked        = "request blocked"
	ErrStrConfig                = "config error"
	ErrStrLoginTokenNotMatch    = "token not match"
	ErrStrCaptcha               = "captcha error"
	ErrStrPaymentMethodNotFound = "payment method not found"
	ErrStrRemote                = "remote channel error"
	ErrStrUnexpectedCall        = "unexpected call"
)
