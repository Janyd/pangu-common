package common

const (

	// CURD 常用业务状态码
	SuccessCode int    = 200
	SuccessMsg  string = "操作成功"

	ParameterVerificationError    int = 400
	ParameterVerificationErrorMsg     = "参数校验错误"

	SystemErrorCode int    = 500
	SystemErrorMsg  string = "系统错误"
)

var (
	ErrorCode = map[int]string{
		SuccessCode: SuccessMsg,
	}
)
