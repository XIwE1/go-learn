package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 中间件的执行流程 —— 以Next为分界线

// c.Next() 之前的代码在handler之前运行。用于设置任务，如**记录开始时间、验证令牌或使用 c.Set() 设置上下文值**

// c.Next() 主动调用链中的下一个处理函数（另一个中间件或最终的路由处理函数）。中间件的执行在此暂停，直到所有下游处理函数完成。

// c.Next() 之后的代码在handler返回后运行。用于清理、记录响应状态或测量延迟。

// 中间件的调用顺序，更广泛的中间件优先执行，全局 > 分组 > 路由

// 中间件 打印日志
func LoggerMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("%s %s", ctx.Request.Method, ctx.Request.URL.Path)
		// 继续下一步
		ctx.Next()
	}
}

// 中间件 - 处理跨域
func CORSMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 假设校验了白名单
		// origin := ctx.GetHeader("Origin")
		// if isOk, err = ...
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")

		// 对 OPTIONS /user/add，Gin 找不到和 OPTIONS 匹配的路由 → 常返回 404
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent) // 204
			return
		}
		ctx.Next()
	}

	// 配置跨域资源共享 CORS ，控制可以向你发送API请求的外部域
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://example.com"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	//   }))
}

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
