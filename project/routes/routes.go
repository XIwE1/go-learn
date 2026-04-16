package routes

import (
	"myproject/user/handler"
	"myproject/user/repository"
	"myproject/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	RegisterUserRoutes(router, userHandler)
	RegisterTest(router)
}
