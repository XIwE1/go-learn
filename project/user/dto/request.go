package dto

import "myproject/user/model"

type UserInfoURI struct {
	Name string `uri:"name" binding:"required"`
	ID   int    `uri:"id" binding:"required,min=1"`
}

type UserListQuery struct {
	Page int    `form:"page,default=1" binding:"min=1"`
	Size int    `form:"size,default=5" binding:"min=1,max=100"`
	Sort string `form:"sort"`
}

type UserCreate struct {
	Name  string `uri:"name" json:"name" binding:"required"`
	Email string `uri:"email" json:"email" binding:"required"`
}

type UserDelete struct {
	Id int `uri:"id" json:"id" binding:"required"`
}

type UserUpdate struct {
	User model.User
}
