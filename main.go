package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"myproject/logs"
	"myproject/routes"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// 中间层打印日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		// 继续下一步
		c.Next()
		// 中间件的执行流程 —— 以Next为分界线

		// c.Next() 之前的代码在路由处理函数之前运行。用于设置任务，如**记录开始时间、验证令牌或使用 c.Set() 设置上下文值**

		// c.Next() 主动调用链中的下一个处理函数（另一个中间件或最终的路由处理函数）。中间件的执行在此暂停，直到所有下游处理函数完成。

		// c.Next() 之后的代码在路由处理函数返回后运行。用于清理、记录响应状态或测量延迟。

		// 中间件的调用顺序，更广泛的中间件优先执行，全局 > 分组 > 路由
	}
}

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

// 中间件（拦截器），功能：预处理，登录授权、验证、分页、耗时统计...
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		// ... check token, session, etc.
		ctx.Set("usersesion", "userid-1")
		ctx.Next() // 放行
		// ctx.Abort() // 阻止
	}
}

// 中间件 处理路由抛出的错误errors
func ErrorHandler() gin.HandlerFunc {
	// 在路由中会遇到多种错误——无效输入、数据库故障、未授权访问或内部 bug。在每个处理函数中单独处理错误会导致重复代码和不一致的响应
	// 采用集中式的错误处理的中间件
	// 在每个请求后运行并检查通过 c.Error(err) 添加到 Gin 上下文中的任何错误来解决这个问题
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}
		// 拿出错误队列的最后一个错误
		err := ctx.Errors.Last().Err
		var appError *AppError
		// 匹配error并赋值
		if errors.As(err, &appError) {
			ctx.JSON(appError.Status, gin.H{
				"success": false,
				"error":   gin.H{"code": appError.Code, "message": appError.Message},
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   gin.H{"code": "INTERNAL", "message": "an unexpected error occurred"},
			})
		}
	}
}

// 中间件 测试路由级中间件 校验id是否为有效的格式
func ValidateIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if len(id) > 5 {
			Fail(ctx, ErrBadRequest.Status, 102, "不合法，id不能大于5位数")
			ctx.Abort()
		}
		ctx.Next()
	}
}

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

// 每个结构体只服务于一个端点
type UserCreate struct {
	Name string `uri:"name" json:"name" binding:"required"`
}
type ListQuery struct {
	Page int    `form:"page,default=1" binding:"min=1"`
	Size int    `form:"size,default=5" binding:"min=1,max=100"`
	Sort string `form:"sort"`
}

type BaseResp[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
	// Message string `json:"message"`
	Error *ErrorInfo `json:"error,omitempty"`
	Meta  *Meta      `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

type Meta struct {
	Page  int    `json:"page,omitempty"`
	Size  int    `json:"size,omitempty"`
	Sort  string `json:"sort,omitempty"`
	Total int    `json:"total,omitempty"`
}

type ListUserData struct {
	List []User `json:"list"`
}

// AppError 代表 api接口错误时返回信息的结构
type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ErrNotFound     = &AppError{Status: http.StatusNotFound, Code: "NOT_FOUND", Message: "resource not found"}
	ErrUnauthorized = &AppError{Status: http.StatusUnauthorized, Code: "UNAUTHORIZED", Message: "authentication required"}
	ErrBadRequest   = &AppError{Status: http.StatusBadRequest, Code: "BAD_REQUEST", Message: "invalid request"}
)

func (e *AppError) Error() string {
	return e.Message
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResp[any]{
		Code: 100,
		Data: data,
	})
}

func Fail(c *gin.Context, status int, code int, message string) {
	c.JSON(status, BaseResp[any]{
		Code:  code,
		Error: &ErrorInfo{Message: message},
	})
}

func main() {
	// 日志写入
	logs.RegisterLog()

	router := gin.Default()

	// 自定义日志格式
	logs.FormatLogs(router)
	// 忽略指定路由与错误码
	logs.IgnoreLogConfig(router)

	// 使用中间件
	router.Use(Logger())
	router.Use(CORSHandler())
	router.Use(ErrorHandler())
	router.Use(logs.RequestIdMiddleWare())

	expectedHost := "localhost:8080"

	// 业务文件 - 注册路由
	routes.RegisterUserRoutes(router)
	routes.RegisterOrderRoutes(router)

	// 设置安全头
	router.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		// 通过禁止页面在 <iframe> 中加载来防止点击劫持
		c.Header("X-Frame-Options", "DENY")
		// 控制浏览器允许加载哪些资源（脚本、样式、图片、字体等）以及从哪些来源
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		// 强制浏览器在指定的 max-age 期间对所有未来请求使用 HTTPS
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// 控制传出请求中发送多少引用者信息
		c.Header("Referrer-Policy", "strict-origin")
		// 防止 MIME 类型嗅探攻击，否则攻击者可能会将有害脚本伪装成无害文件
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	// 配置跨域资源共享 CORS ，控制可以向你发送API请求的外部域
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://example.com"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	//   }))

	// 对于 需要共享的 数据库连接、配置和服务等依赖 ，GO推荐了三种模式（而不是DI）
	// **首先 连接数据库**
	// db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// 1. 闭包模式 - 适合中小型项目 少量依赖 - 依赖通过闭包捕获进 Handler
	// func XXHandler(db *sql.DB) gin.HandlerFunc { return func(c *gin.Context) {}} ...
	// r.GET("/ping", PingHandler(db))
	// r.GET("/users/:id", GetUserHandler(db))

	// 2. 结构体 + 方法模式 - 中大型应用 - 编译期类型安全
	// type App struct { DB     *sql.DB, Logger *slog.Logger }
	// app := &App{ DB:     db,}
	// r.GET("/users/:id", app.GetUser)

	// 3. 中间件注入 - AOP - 通过在上下文注入依赖
	// func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	// 	return func(c *gin.Context) {
	// 	  c.Set("db", db)
	// 	  c.Next()
	// 	}
	//   }
	//   func GetUser(c *gin.Context) {
	// 	db := c.MustGet("db").(*sql.DB)
	// 	// Use db...
	//   }
	// r.Use(DatabaseMiddleware(db))
	// r.GET("/users/:id", GetUser)
	// 优雅地处理连接错误
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	// 不管使用 database/sql 还是 GORM
	// 都要将请求上下文传递给查询。 使用 c.Request.Context()，以便在客户端断开连接或超时触发时取消长时间运行的查询
	// 当请求需要执行多个必须一起成功或失败的写操作时，使用**数据库事务**

	// **会话管理** 当你需要跨微服务的无状态认证时使用 JWT
	// 创建一个存储session会话的地方
	// store := cookie.NewStore([]byte("your-secret-key"))
	// r.Use(sessions.Sessions("mysession", store))
	// r.GET("/login", func(c *gin.Context) {
	// 	session := sessions.Default(c)
	// 	session.Set("user", "john")
	// 	session.Save()
	// 	user := session.Get("user")
	// 	session.Clear()
	// 	session.Save()
	// 	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
	// })

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

		c.JSON(http.StatusOK,
			// √ 定义内层结构体和外层泛型
			BaseResp[ListUserData]{
				Code: 0,
				Data: ListUserData{
					List: db_list,
				},
				Meta: &Meta{
					Page:  query.Page,
					Size:  query.Size,
					Sort:  query.Sort,
					Total: 100,
				},
			},

			// × 零散拼接 没有约束和复用性
			// 	gin.H{
			// 	"code":    0,
			// 	"message": "success",
			// 	"data": gin.H{
			// 		"list":  db_list,
			// 		"page":  query.Page,
			// 		"size":  query.Size,
			// 		"sort":  query.Sort,
			// 		"total": 100,
			// 	},
			// }
		)
	})

	router.POST("/user/add", func(c *gin.Context) {
		var user UserCreate
		err := c.ShouldBindJSON(&user)
		// name := c.DefaultPostForm("name", "nobody")

		if err != nil {
			c.JSON(http.StatusBadRequest, BaseResp[any]{
				Code:  100,
				Error: &ErrorInfo{Message: "创建失败"},
			})
			return
		}

		// 数据库模拟添加一条数据
		// newUser := db.CreateUser(&user)

		c.JSON(http.StatusOK, BaseResp[User]{
			Code: 0,
			Data: User{
				Name: user.Name,
				Id:   123,
			},
		})
	})

	// **路由分组**
	{
		v1 := router.Group("/v1")
		v1.POST("/login", func(ctx *gin.Context) {})
		v1.POST("/submit", func(ctx *gin.Context) {})

		v2 := router.Group("/v2")
		v2.POST("/login", func(ctx *gin.Context) {})
		v2.POST("/submit", func(ctx *gin.Context) {})
	}

	// **分组使用中间件**
	// public routes -- 不需要权限校验
	public := router.Group("/api")
	{
		public.GET("/health", func(ctx *gin.Context) {})
		// **嵌套分组**
		users := public.Group("/v1")
		// /api/v1/users
		users.GET("/users", func(ctx *gin.Context) {})
	}
	// private routes -- 需要权限校验
	private := router.Group("/api")
	private.Use(AuthRequired())
	{
		private.POST("/settings", func(ctx *gin.Context) {})
	}

	// 测试一致响应格式
	router.GET("/api/user/:id", ValidateIdHandler(), func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "0" {
			Fail(ctx, http.StatusNotFound, 201, "invalid id")
			return
		}
		Ok(ctx, gin.H{"name": "xiwei", "id": id})
	})

	router.GET("/api/articles", func(ctx *gin.Context) {
		cursor := ctx.DefaultQuery("cursor", "")
		limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

		if limit > 100 {
			limit = 100
		}

		// articles, nextCursor := db.ListArticles(cursor, limit)
		_ = cursor

		Ok(ctx, gin.H{
			"code":        100,
			"data":        []gin.H{}, // articles
			"next_cursor": "",        // nextCursor (empty string means no more pages)
		})
	})

	// 筛选参数demo
	// GET /api/products?category=electronics&min_price=10&sort=price&order=asc
	router.GET("/api/products", func(ctx *gin.Context) {
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

	// 测试自定义错误处理
	router.GET("/api/items/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "0" {
			// 将错误信息添加到上下文context的错误列表中
			ctx.Error(ErrNotFound)
			return
		}

		Ok(ctx, gin.H{"success": true, "data": gin.H{"id": id}})
	})

	// gin中的goroutine
	router.GET("/long_async/:id", func(c *gin.Context) {
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
