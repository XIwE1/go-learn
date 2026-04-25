package dto

type UserInfoURI struct {
	Name string `uri:"name" binding:"required"`
	ID   int    `uri:"id" binding:"required,min=1"`
}

type UserListQuery struct {
	Name   string `form:"name" json:"name"`
	Email  string `form:"email" json:"email"`
	Page   int    `form:"page,default=1" binding:"min=1"`
	Size   int    `form:"size,default=5" binding:"min=1,max=100"`
	Sort   string `form:"sort"`
	SortBy string `form:"sort_by"`
}

type UserCreate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserDelete struct {
	Id uint `json:"id" binding:"required"`
}

type UserUpdate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Id    uint   `json:"id" binding:"required,min=1"`
}
