package usecase

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/cacheclient"
	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/Akito-Fujihara/web-application-template/app/domain/repository"
	"github.com/google/uuid"
)

type IAccountUsecase interface {
	AuthenticateUser(ctx context.Context, email string, password string) (sessionID string, err error)
	CreateUser(ctx context.Context, user *model.User, password string) error
}

type AccountUsecase struct {
	cacheClient cacheclient.ICacheClient
	userRepo repository.IUserRepository
}

func NewAccountUsecase(
	cacheClient cacheclient.ICacheClient,
	userRepo repository.IUserRepository,
) IAccountUsecase {
	return &AccountUsecase{
		cacheClient: cacheClient,
		userRepo: userRepo,
	}
}

func (u *AccountUsecase) AuthenticateUser(ctx context.Context, email string, password string) (sessionID string, err error) {
	user, err := u.userRepo.ValidateCredentials(ctx, email, password)
	if err != nil {
		return "", err
	} 
	sessionID = uuid.New().String()
	if err := u.cacheClient.SetSession(ctx, sessionID, user); err != nil {
		return "", err
	}
	return sessionID, nil
}

func (u *AccountUsecase) CreateUser(ctx context.Context, user *model.User, password string) error {
	return u.userRepo.CreateUser(ctx, user, password)
}
