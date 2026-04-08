package main

import (
	"myproject/routes"
	"myproject/user/handler"
	"myproject/user/service"
	"net/http"
	"time"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// 中间件 - 处理跨域
func CORSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设校验了白名单
		// origin := c.GetHeader("Origin")
		// if isOk, err = ...
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		// 对 OPTIONS /user/add，Gin 找不到和 OPTIONS 匹配的路由 → 常返回 404
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // 204
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSHandler())

	// 注册User相关的路由
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(router, userHandler)

	// 通过 http.Cookie 设置 cookie
	router.GET("/getCookie", func(ctx *gin.Context) {
		ctx.SetCookieData(&http.Cookie{
			Name:     "session_id",
			Value:    "abc123",
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(24 * time.Hour),
			MaxAge:   86400,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		ctx.String(http.StatusOK, "ok")
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
