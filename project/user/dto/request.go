package dto

type UserInfoURI struct {
	Name string `uri:"name" binding:"required"`
	ID   int    `uri:"id" binding:"required,min=1"`
}

type UserListQuery struct {
	Page int    `form:"page,default=1" binding:"min=1"`
	Size int    `form:"size,default=5" binding:"min=1,max=100"`
	Sort string `form:"sort"`
}
