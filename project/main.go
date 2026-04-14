package main

import (
	"io"
	"log/slog"
	"myproject/common/log"
	"myproject/middleware"
	"myproject/routes"
	"net/http"
	"os"
	"time"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// 绑定中间件
	router.Use(gin.Recovery())
	logFile := log.InitLogWriter()
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), nil))
	middleware.RegisterGlobal(router, logger)

	// 注册路由
	routes.Register(router)

	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 默认启动方式
	// router.Run() // listens on 0.0.0.0:8080 by default

	// 使用自定义服务器配置
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
