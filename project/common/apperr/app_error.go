package apperr

import "net/http"

// AppError 代表 api接口错误时返回信息的结构
type AppError struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrNotFound     = &AppError{Status: http.StatusNotFound, Code: 404, Message: "resource not found"}
	ErrUnauthorized = &AppError{Status: http.StatusUnauthorized, Code: 401, Message: "authentication required"}
	ErrBadRequest   = &AppError{Status: http.StatusBadRequest, Code: 400, Message: "invalid request"}
)

func (e *AppError) Error() string {
	return e.Message
}
