package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var SkipPaths = []string{"/ping"}

// 中间件 - 结构化日志记录
func SlogMiddleWare(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		requestID, _ := ctx.Get("request_id")

		logger.Info("request",
			slog.String("request_id", requestID.(string)),
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.Int("status", ctx.Writer.Status()),
			slog.Duration("latency", time.Since(start)),
			slog.String("client_ip", ctx.ClientIP()),
		)

		if len(ctx.Errors) > 0 {
			for _, err := range ctx.Errors {
				logger.Error("request error", slog.String("error", err.Error()))
			}
		}
	}
}

// 中间件 自定义 log 的格式
func FormatLogMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// 中间件 忽略指定路由的记录
func IgnoreLogMiddleware() gin.HandlerFunc {
	// 直接设置 跳过指定路由的日志记录
	loggerConfig := gin.LoggerConfig{SkipPaths: SkipPaths}

	// 设置函数 判断要跳过的记录
	loggerConfig.Skip = func(c *gin.Context) bool {
		// as an example skip non server side errors
		return c.Writer.Status() < http.StatusInternalServerError
	}
	return gin.LoggerWithConfig(loggerConfig)
}
