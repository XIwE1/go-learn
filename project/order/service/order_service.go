package service

import (
	"errors"
	"myproject/user/dto"
)

var ErrInvalidUser = errors.New("invalid data")

// 定义userService有哪些方法
type UserService interface {
	GetUserInfo(name string, id int) (dto.UserInfoResp, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) GetUserInfo(name string, id int) (dto.UserInfoResp, error) {
	if name == "" || id <= 0 {
		return dto.UserInfoResp{}, ErrInvalidUser
	}

	// TODO: 数据库查询返回
	return dto.UserInfoResp{
		Name: name,
		ID:   id,
	}, nil
}
