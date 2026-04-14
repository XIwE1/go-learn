package routes

import (
	"log"
	"myproject/common/apperr"
	"myproject/common/httpx"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterTest(router *gin.Engine) {
	test := router.Group("/test")

	// gin中的goroutine
	test.GET("/long_async/:id", func(c *gin.Context) {
		// Gin 为了性能会用 sync.Pool 复用 **gin.Contex**。
		// 请求的 handler 返回后，这个 Context 可能被放回池里、很快又被另一个请求复用

		// 使用copy创建一个快照 避免被Pool池影响
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)

			// 如果你的 goroutine 还握着旧的 c 在读写，**就可能读到/写到“另一个请求”的上下文**，导致竞态、数据错乱甚至 panic
			// log.Println("Done! in path " + c.Request.URL.Path)
		}()
	})

	// 通过 http.Cookie 设置 cookie
	test.GET("/getCookie", func(ctx *gin.Context) {
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

	// 测试自定义错误处理
	test.GET("/error/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "0" {
			// 将错误信息添加到上下文context的错误列表中
			ctx.Error(apperr.ErrNotFound)
			return
		}

		httpx.Ok(ctx, gin.H{"id": id, "success": true})
	})

	// 筛选参数demo
	// GET /test/products?category=electronics&min_price=10&sort=price&order=asc
	test.GET("/products", func(ctx *gin.Context) {
		category := ctx.Query("category")
		minPrice := ctx.DefaultQuery("min_price", "0")
		maxPrice := ctx.DefaultQuery("max_price", "9999")
		order := ctx.DefaultQuery("order", "asc")
		sortBy := ctx.DefaultQuery("sort", "created_at")

		// 校验排序字段 以防代码注入
		allowed := map[string]bool{"created_at": true, "price": true, "name": false}
		if !allowed[sortBy] {
			sortBy = "created_at"
		}
		if order != "desc" && order != "asc" {
			order = "desc"
		}

		// 通过传递来的字段执行一些查询操作
		_ = category
		_ = minPrice
		_ = maxPrice

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []gin.H{},
			"filters": gin.H{
				"category":  category,
				"min_price": minPrice,
				"max_price": maxPrice,
				"sort":      sortBy,
				"order":     order,
			},
		})
	})
}
