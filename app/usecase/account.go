package usecase

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/Akito-Fujihara/web-application-template/app/domain/repository"
)

type IAccountUsecase interface {
	CreateUser(ctx context.Context, user *model.User, password string) error
}

type AccountUsecase struct {
	userRepo repository.IUserRepository
}

func NewAccountUsecase(userRepo repository.IUserRepository) IAccountUsecase {
	return &AccountUsecase{
		userRepo: userRepo,
	}
}

func (u *AccountUsecase) CreateUser(ctx context.Context, user *model.User, password string) error {
	return u.userRepo.CreateUser(ctx, user, password)
}
