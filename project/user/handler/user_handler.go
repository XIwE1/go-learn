package handler

import (
	"errors"
	"myproject/user/dto"
	"myproject/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetUserInfo(ctx *gin.Context) {
	var req dto.UserInfoURI
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "路径参数不合法",
			"error":   err.Error(),
		})
		return
	}

	resp, err := handler.service.GetUserInfo(req.Name, req.ID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUser) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
