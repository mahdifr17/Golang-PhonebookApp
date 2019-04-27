package structs

// BaseResponse is a base struct to construct error response
type BaseResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}
