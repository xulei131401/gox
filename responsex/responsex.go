package responsex

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewApiResponse(code int, msg string, data interface{}) *ApiResponse {
	return &ApiResponse{Code: code, Msg: msg, Data: data}
}

func NewSuccessApiResponse() *ApiResponse {
	return NewApiResponse(SuccessCode, "success", nil)
}

func NewSuccessApiResponseWithData(data interface{}) *ApiResponse {
	return NewApiResponse(SuccessCode, "success", data)
}
