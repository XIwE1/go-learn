package routes

import (
	"myproject/user/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler) {
	users := router.Group("/user")

	users.GET("/info/:name/:id", handler.GetUserInfo)
	users.GET("/list", handler.GetUserList)
	users.POST("/add", handler.CreateUser)
	users.DELETE("/delete", handler.DeleteUser)
	users.POST("/update", handler.UpdateUser)
}
