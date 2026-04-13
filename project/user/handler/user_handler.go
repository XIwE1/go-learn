package handler

import (
	"errors"
	"myproject/common/apperr"
	"myproject/common/httpx"
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
		httpx.FailApp(ctx, apperr.ErrBadRequest)
		return
	}

	resp, err := handler.service.GetUserInfo(req.Name, req.ID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUser) {
			httpx.FailApp(ctx, apperr.ErrBadRequest)
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
