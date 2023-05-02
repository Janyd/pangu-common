package common

func ErrMsg(code int) string {
	msg, ok := ErrorCode[code]
	if !ok {
		msg = ErrorCode[SystemErrorCode]
	}
	return msg
}

// MessageError 自定义消息错误
type MessageError struct {
	err       error
	code      int
	msg       string
	secondMsg string
}

// 集中判断常见错误
func Error(err error) *MessageError {
	return New(SystemErrorCode)
}

func New(code int) *MessageError {
	return NewMessageError(nil, code)
}

func NewMessageError(parent error, code int) *MessageError {
	return &MessageError{
		err:  parent,
		code: code,
		msg:  ErrMsg(code),
	}
}

func NewSecondMsg(parent error, code int, msg string) *MessageError {
	return &MessageError{
		err:       parent,
		code:      code,
		msg:       ErrMsg(code),
		secondMsg: msg,
	}
}

func (m *MessageError) Code() int {
	return m.code
}

func (m *MessageError) Error() string {
	return ErrorCode[m.code]
}

func (m *MessageError) Parent() error {
	return m.err
}

func (m *MessageError) SecondMsg() string {
	return m.secondMsg
}
