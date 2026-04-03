package logs

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterLog() {
	f, _ := os.Create("gin.log")
	// 会覆盖终端打印
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果想打印日志的同时记录日志 使用这个
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
