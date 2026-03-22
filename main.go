package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 中间层打印日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		// 继续下一步
		c.Next()
	}
}

func CORSHandler() gin.HandlerFunc {
	// 假设校验了白名单
	// if isOk, err = ...
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// 中间件（拦截器），功能：预处理，登录授权、验证、分页、耗时统计...
// func myHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
// 		ctx.Set("usersesion", "userid-1")
// 		ctx.Next()  // 放行
// 		ctx.Abort() // 阻止
// 	}
// }

func main() {
	router := gin.Default()
	// 使用中间件

	router.Use(Logger())
	router.Use(CORSHandler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
