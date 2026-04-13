package dto

import (
	"myproject/common/response"
	"myproject/user/model"
)

type UserInfoResp struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type UserListResp struct {
	List []model.User  `json:"userList"`
	Meta response.Meta `json:"meta"`
}
