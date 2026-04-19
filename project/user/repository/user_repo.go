package repository

import (
	"context"
	"myproject/user/dto"
	"myproject/user/model"

	"gorm.io/gorm"
)

// 定义抽象接口 给 userRepository 结构体去实现
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) (model.User, error)
	Search(ctx context.Context, meta dto.UserListQuery) ([]model.User, int64, error)
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

func (userRepo *userRepository) Delete(ctx context.Context, id uint) (model.User, error) {
	// _, err := gorm.G[model.User](userRepo.db).Where("id = ?", id).Delete(ctx)

	// 返回被删除的数据
	var deleted model.User
	err := userRepo.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			if err := tx.First(&deleted, id).Error; err != nil {
				return err
			}
			return tx.Delete(&deleted).Error
		})
	if err != nil {
		return model.User{}, err
	}
	return deleted, err
}

func (userRepo *userRepository) Search(ctx context.Context, meta dto.UserListQuery) ([]model.User, int64, error) {
	offset := (meta.Page - 1) * meta.Size
	g := gorm.G[model.User](userRepo.db)
	total, err := g.Count(ctx, "*")
	if err != nil {
		return nil, 0, err
	}

	users, err := g.Offset(offset).Limit(meta.Size).Find(ctx)

	return users, total, err
}
