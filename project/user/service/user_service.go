package service

import (
	"errors"
	"fmt"
	"myproject/common/response"
	"myproject/user/dto"
	"myproject/user/model"
)

var ErrInvalidUser = errors.New("invalid data")

// 定义userService有哪些方法
type UserService interface {
	GetUserInfo(name string, id int) (dto.UserInfoResp, error)
	GetUserList(query dto.UserListQuery) (dto.UserListResp, error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

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

func (us *userService) GetUserList(query dto.UserListQuery) (dto.UserListResp, error) {
	// 模拟从数据库search到对应数据
	db_list := make([]model.User, 0, query.Size)
	for i := 0; i < query.Size; i++ {
		db_list = append(db_list, model.User{
			Name: fmt.Sprintf("user-%d", i+1),
			Id:   i + 1,
		})
	}

	return dto.UserListResp{
		List: db_list,
		Meta: response.Meta{
			Page:  query.Page,
			Size:  query.Size,
			Sort:  query.Sort,
			Total: 100,
		},
	}, nil
}
