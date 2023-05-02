package common

// Result 统一返回结果
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResError(bizerr ResultError) *Result {
	return &Result{Code: bizerr.Code(), Msg: bizerr.Error()}
}

func ServerError() *Result {
	return ResError(ServerErrorResult)
}

func Success(data interface{}) *Result {
	return &Result{Code: SuccessCode, Msg: SuccessMsg, Data: data}
}
