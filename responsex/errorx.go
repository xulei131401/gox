package responsex

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultCodeError(msg string) error {
	return NewCodeError(FailCode, msg)
}

// go-zero全局错误处理器不能返回实现Error接口的结构体，否则以文本处理
type codeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeErrorResponse(code int, msg string) *codeErrorResponse {
	return &codeErrorResponse{Code: code, Msg: msg}
}

func (e *CodeError) Data() *codeErrorResponse {
	return NewCodeErrorResponse(e.Code, e.Msg)
}
