package middleware

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

var IgnorePaths = []string{"/ping"}

// 中间件 控制台打印日志
func LoggerMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("%s %s", ctx.Request.Method, ctx.Request.URL.Path)
		// 继续下一步
		ctx.Next()
	}
}

// 中间件 - 结构化日志记录
func SlogMiddleWare(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		// if shouldSkip(ctx) {
		// 	return
		// }

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
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
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
		},
		// SkipPaths: IgnorePaths,
		// Skip: shouldSkip,
	})
}

// 判断是否跳过的函数
func shouldSkip(context *gin.Context) bool {
	param := context.Request.URL.Path

	if slices.Contains(IgnorePaths, param) {
		return true
	}

	return context.Writer.Status() < http.StatusInternalServerError
}
