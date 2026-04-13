package httpx

import (
	"myproject/common/apperr"
	"myproject/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 成功响应
func Ok(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, response.BaseResp[any]{
		Code: 0,
		Data: data,
	})
}

// 失败响应
func Fail(ctx *gin.Context, status int, code int, message string) {
	ctx.JSON(status, response.BaseResp[any]{
		Code:  code,
		Error: &response.ErrorInfo{Message: message},
	})
}

// 指定格式的失败响应
func FailApp(ctx *gin.Context, err *apperr.AppError) {
	Fail(ctx, err.Status, err.Code, err.Message)
}
