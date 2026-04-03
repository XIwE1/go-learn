package logs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

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

func RegisterLog() {
	f, _ := os.Create("gin.log")
	// 会覆盖终端打印
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果想打印日志的同时记录日志 使用这个
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
