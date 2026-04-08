package dto

type UserInfoURI struct {
	Name string `uri:"name" binding:"required"`
	ID   int    `uri:"id" binding:"required,min=1"`
}
