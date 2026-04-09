package middleware

import (
	"errors"
	"myproject/common/apperr"
	"myproject/common/httpx"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 中间件 集中式的错误处理
func ErrorMiddleWare() gin.HandlerFunc {
	// 在路由中会遇到多种错误——无效输入、数据库故障、未授权访问或内部 bug。在每个处理函数中单独处理错误会导致重复代码和不一致的响应
	// 集中处理路由抛出的错误errors
	// 在每个请求后运行并检查通过 c.Error(err) 添加到 Gin 上下文中的任何错误来解决这个问题
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}
		// 拿出错误队列的最后一个错误
		err := ctx.Errors.Last().Err
		var appError *apperr.AppError
		// 匹配error并赋值
		if errors.As(err, &appError) {
			httpx.Fail(ctx, appError.Status, appError.Code, appError.Message)
		} else {
			httpx.Fail(ctx, http.StatusInternalServerError, 500, "an unexpected error occurred")
		}
	}
}
