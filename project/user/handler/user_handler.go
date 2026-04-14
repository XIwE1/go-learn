package handler

import (
	"errors"
	"myproject/common/apperr"
	"myproject/common/httpx"
	"myproject/user/dto"
	"myproject/user/service"

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
		httpx.FailApp(ctx, apperr.ErrInternal)
		return
	}

	httpx.Ok(ctx, resp)
}

func (handler *UserHandler) GetUserList(ctx *gin.Context) {
	var query dto.UserListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httpx.FailApp(ctx, apperr.ErrBadRequest)
		return
	}

	resp, _ := handler.service.GetUserList(query)
	httpx.Ok(ctx, resp)
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var user dto.UserCreate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpx.FailApp(ctx, apperr.ErrBadRequest)
		return
	}

	resp, _ := handler.service.CreateUser(user)
	httpx.Ok(ctx, resp)
}
