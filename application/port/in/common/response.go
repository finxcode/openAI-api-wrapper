package common

type Response struct {
	ErrCode int         `json:"err_code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
