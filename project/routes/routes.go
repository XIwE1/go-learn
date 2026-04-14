package routes

import (
	"myproject/user/handler"
	"myproject/user/service"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	RegisterUserRoutes(router, userHandler)
	RegisterTest(router)
}
