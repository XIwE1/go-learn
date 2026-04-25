package service

import (
	"context"
	"errors"
	"myproject/common/response"
	"myproject/user/dto"
	"myproject/user/model"
	"myproject/user/repository"
)

var ErrInvalidUser = errors.New("invalid data")

// 定义userService有哪些方法
type UserService interface {
	GetUserInfo(name string, id int) (dto.UserInfoResp, error)
	GetUserList(ctx context.Context, query dto.UserListQuery) (dto.UserListResp, error)
	CreateUser(ctx context.Context, params dto.UserCreate) (model.User, error)
	DeleteUser(ctx context.Context, info dto.UserDelete) (model.User, error)
	UpdateUser(ctx context.Context, data dto.UserUpdate) (model.User, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

type userService struct {
	userRepo repository.UserRepository
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

func (us *userService) GetUserList(ctx context.Context, query dto.UserListQuery) (dto.UserListResp, error) {
	// 模拟从数据库search到对应数据
	// db_list := make([]model.User, 0, query.Size)
	// for i := 0; i < query.Size; i++ {
	// 	db_list = append(db_list, model.User{
	// 		Name: fmt.Sprintf("user-%d", i+1),
	// 	})
	// }

	db_list, total, err := us.userRepo.Search(ctx, query)

	return dto.UserListResp{
		List: db_list,
		Meta: response.Meta{
			Page:   query.Page,
			Size:   query.Size,
			Sort:   query.Sort,
			SortBy: query.SortBy,
			Total:  total,
		},
	}, err
}

func (us *userService) CreateUser(ctx context.Context, params dto.UserCreate) (model.User, error) {
	// 数据库模拟添加一条数据
	// newUser := db.CreateUser(&user)

	// newUser := model.User{
	// 	Name: params.Name,
	// 	Id:   uuid.New().ClockSequence(),
	// }

	newUser := &model.User{
		Name:  params.Name,
		Email: params.Email,
	}

	if err := us.userRepo.Create(ctx, newUser); err != nil {
		return model.User{}, err
	}

	return *newUser, nil
}

func (us *userService) DeleteUser(ctx context.Context, info dto.UserDelete) (model.User, error) {
	// 数据库模拟删除一条数据
	// targetUser := db.DeleteUser(&user)

	// targetUser := model.User{
	// 	Name: uuid.New().String(),
	// }

	targetUser, err := us.userRepo.Delete(ctx, info.Id)

	return targetUser, err
}

func (us *userService) UpdateUser(ctx context.Context, data dto.UserUpdate) (model.User, error) {
	// 数据库模拟更新一条数据
	// updatedUser := db.UpdateUser(&user)

	// updatedUser := model.User{
	// 	Name: data.User.Name,
	// }

	updatedUser, err := us.userRepo.Update(ctx, data)

	return updatedUser, err
}
