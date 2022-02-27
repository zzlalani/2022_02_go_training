package services

import (
	"github.com/zzlalani/go-practice/internal/dto"
	"github.com/zzlalani/go-practice/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo,
	}
}

func (u *UserService) CreateUser(dto dto.UserReq) (uint, error) {
	user := repository.User{
		Username: dto.User,
		Password: dto.Pass,
	}

	return u.userRepo.Insert(&user)
}
