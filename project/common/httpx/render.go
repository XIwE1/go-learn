package httpx

import (
	"myproject/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 成功响应
func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response.BaseResp[any]{
		Code: 0,
		Data: data,
	})
}

// 失败响应
func Fail(c *gin.Context, status int, code int, message string) {
	c.JSON(status, response.BaseResp[any]{
		Code:  code,
		Error: &response.ErrorInfo{Message: message},
	})
}
