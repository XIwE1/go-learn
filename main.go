package main

import (
	"fmt"
	"log"
	"net/http"

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

type User struct {
	// uri 结构体标签将 URI 路径参数直接绑定到结构体中
	Name string `uri:"name" json:"name" binding:"required"`
	Id   int    `uri:"id" json:"id" binding:"required"`

	// `xx:"yy"` = 结构体标签（Struct Tag）。它是一种元数据（关于数据的数据），用来为结构体的字段提供额外的信息
	// Name string `json:"name"`
	// 当把 User 结构体转换成 JSON 字符串（序列化）时，字段 Name 在 JSON 中应该使用键名 "name"，而不是默认的字段名 "Name"。
	// 类似的还有
	// gorm:"column:user_name;type:varchar(100)" 用来指定数据库表中的列名和类型
}

type ListQuery struct {
	Page int    `form:"page,default=1" binding:"min=1"`
	Size int    `form:"size,default=5" binding:"min=1,max=100"`
	Sort string `form:"sort"`
}

func main() {
	router := gin.Default()
	// 使用中间件

	router.Use(Logger())
	router.Use(CORSHandler())
	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 查询用户接口 路径参数
	router.GET("/user/info/:name/:id", func(c *gin.Context) {
		var user User
		// √ 使用 Gin 的绑定机制 将请求参数直接绑定到结构体，并**自动进行类型转换和校验**
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "路径参数不合法",
				"error":   err.Error(),
			})
			return
		}
		// 在处理函数内使用 c.Param("name") 来获取路径参数的值
		userName := user.Name
		userId := user.Id
		if userName == "" || userId <= 0 {
			c.JSON(http.StatusBadRequest, "无效的数据")
			return
		}

		// × 使用 strconv.Atoi 最基础的方法，但需要在每个 handler 中重复编写错误处理代码
		// id, err := strconv.Atoi(idStr)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		// 	return
		// }

		c.JSON(http.StatusOK, user)
	})

	// 查询表格接口 字符串参数
	router.GET("/list", func(c *gin.Context) {
		// page := c.DefaultQuery("page", "1")
		// _page, err := strconv.Atoi(page)
		// if err ...
		// size := c.DefaultQuery("size", "5")
		// sort := c.Query("sort")

		var query ListQuery
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1001,
				"data":    nil,
				"message": "参数不合法",
			})
			return
		}

		// 模拟从数据库search到对应数据
		db_list := make([]User, 0, query.Size)
		for i := 0; i < query.Size; i++ {
			db_list = append(db_list, User{
				Name: fmt.Sprintf("user-%d", i+1),
				Id:   i + 1,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data": gin.H{
				"list":  db_list,
				"page":  query.Page,
				"size":  query.Size,
				"sort":  query.Sort,
				"total": 100,
			},
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
