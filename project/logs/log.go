package logs

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var SkipPaths = []string{"/ping"}

// 中间件 - 生成 requestId 以便追踪请求
func RequestIdMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-ID")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx.Set("request_id", requestId)
		ctx.Header("X-Request-ID", requestId)
		ctx.Next()
	}
}

// 中间件 - 结构化日志记录
func SlogMiddleWare(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		requestID, _ := c.Get("request_id")

		logger.Info("request",
			slog.String("request_id", requestID.(string)),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("latency", time.Since(start)),
			slog.String("client_ip", c.ClientIP()),
		)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error("request error", slog.String("error", err.Error()))
			}
		}
	}
}

// 自定义 log 的格式
func FormatLogs(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
	}))
}

// 配置忽略的路由和状态码
func IgnoreLogConfig(router *gin.Engine) {
	// 跳过指定路由的日志记录
	loggerConfig := gin.LoggerConfig{SkipPaths: SkipPaths}

	// 可以设置函数判断要跳过的记录
	loggerConfig.Skip = func(c *gin.Context) bool {
		// as an example skip non server side errors
		return c.Writer.Status() < http.StatusInternalServerError
	}
	router.Use(gin.LoggerWithConfig(loggerConfig))
}

func RegisterLog() *os.File {
	f, _ := os.Create("gin.log")
	// 会覆盖终端打印
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果想打印日志的同时记录日志 使用这个
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 返回 os.File 让其他插件也能写入同一个文件
	return f
}
