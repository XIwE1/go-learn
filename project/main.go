package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	commonLog "myproject/common/log"
	"myproject/middleware"
	"myproject/routes"
	"myproject/user/model"
	"net/http"
	"os"
	"time"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "developer_1:txyprtwy12AA!@tcp(115.190.227.247:3306)/gin_mysql?charset=utf8mb4&parseTime=True&loc=Local"

	// 可配置：从环境变量中获取dsn
	// dsn := os.Getenv("DATABASE_DSN")
	// if dsn == "" {
	// 	dsn = os.Getenv("MYSQL_DSN")
	// }
	// if dsn == "" {
	// 	log.Fatal("请设置环境变量 DATABASE_DSN 或 MYSQL_DSN（MySQL DSN 字符串）")
	// }
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	fmt.Println("db = ", db)
	fmt.Println("err = ", err)
	// 创建对应的表
	db.AutoMigrate(&model.User{})

	// 连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	router := gin.New()

	// 绑定中间件
	router.Use(gin.Recovery())
	logFile := commonLog.InitLogWriter()
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), nil))
	middleware.RegisterGlobal(router, logger)

	// 注册路由
	routes.Register(router, db)

	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 默认启动方式
	// router.Run() // listens on 0.0.0.0:8080 by default

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// 使用自定义服务器配置
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
