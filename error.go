package common

type ResultError struct {
	code int
	err  string
}

var (
	SuccessResult        = NewResultError(200, "success")
	BizErrResult         = NewResultError(400, "biz error")
	MethodNotAllowResult = NewResultError(405, "method not allow")
	NotFoundResult       = NewResultError(404, "not found")
	ServerErrorResult    = NewResultError(500, "server error")
)

func (e ResultError) Error() string {
	return e.err
}

func (e ResultError) Code() int {
	return e.code
}

// NewResultError 创建业务逻辑错误结构体，默认为业务逻辑错误
func NewResultError(code int, msg string) ResultError {
	return ResultError{
		code: code,
		err:  msg,
	}
}
