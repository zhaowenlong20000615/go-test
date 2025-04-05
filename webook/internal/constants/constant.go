package constants

import "errors"

var (
	ErrDuplicateEmail        = errors.New("该邮箱已被注册")
	ErrNotFoundUser          = errors.New("该用户不存在")
	ErrInvaildUserOrPassword = errors.New("邮箱或密码不对")
)

const (
	USER_ID            = "userId"
	SHORT_TIME_JWT_KEY = "shortTimeJwtKey"
	LONG_TIME_JWT_KEY  = "longTimeJwtKey"
	CLAIMS_KEY         = "claims"
)
