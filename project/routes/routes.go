package routes

import (
	"myproject/order/service"
	"myproject/user/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	RegisterUserRoutes(router, userHandler)

}
