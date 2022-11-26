package services

import (
	"context"
	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type UserService interface {
	GetAllUser() ([]dto.User, common.SubReturnCode)
	GetUserByUsername(ctx context.Context, username string) (dto.User, common.SubReturnCode)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func (u *userServiceImpl) GetUserByUsername(ctx context.Context, username string) (dto.User, common.SubReturnCode) {
	val, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		glog.Infoln("GetAllUser service err: ", err)
		return dto.User{}, common.SystemError
	}
	return dto.User{
		ID:         val.ID,
		UserName:   val.UserName,
		CreatedAt:  val.CreatedAt,
		UpdatedAt:  val.UpdatedAt,
		DeletedAt:  val.DeletedAt,
		PassWord:   val.PassWord,
		IsResolver: val.IsResolver,
		Email:      val.Email,
		Phone:      val.Phone,
	}, common.OK
}

func (u *userServiceImpl) GetAllUser() ([]dto.User, common.SubReturnCode) {
	var resp = make([]dto.User, 0)
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		glog.Infoln("GetAllUser service err: ", err)
		return resp, common.SystemError
	}
	for _, val := range users {
		resp = append(resp, dto.User{
			ID:         val.ID,
			UserName:   val.UserName,
			CreatedAt:  val.CreatedAt,
			UpdatedAt:  val.UpdatedAt,
			DeletedAt:  val.DeletedAt,
			PassWord:   val.PassWord,
			IsResolver: val.IsResolver,
			Email:      val.Email,
			Phone:      val.Phone,
		})
	}
	return resp, common.OK
}