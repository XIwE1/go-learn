package logs

import (
	"fmt"
	"io"
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
		ctx.Set("request-id", requestId)
		ctx.Header("X-Request-ID", requestId)
		ctx.Next()
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

func RegisterLog() {
	f, _ := os.Create("gin.log")
	// 会覆盖终端打印
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果想打印日志的同时记录日志 使用这个
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
