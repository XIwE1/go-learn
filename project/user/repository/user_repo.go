package repository

import (
	"context"
	"myproject/user/model"

	"gorm.io/gorm"
)

// 定义接口 给 userRepository 结构体去实现
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) Create(ctx context.Context, user *model.User) error {
	// traditional API
	// return userRepo.db.WithContext(ctx).Create(u).Error

	// generics API
	return gorm.G[model.User](userRepo.db).Create(ctx, user)
}
