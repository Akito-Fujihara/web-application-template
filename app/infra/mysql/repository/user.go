package repository

import (
	"context"
	"fmt"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	domainRepo "github.com/Akito-Fujihara/web-application-template/app/domain/repository"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/dbschema"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/ormgen"
)

type UserRepositoy struct {}

func NewUserRepository() domainRepo.IUserRepository {
	return &UserRepositoy{}
}

func (r *UserRepositoy) CreateUser(ctx context.Context, user *model.User, password string) error {
	dto := buildUserToDto(user)
	// sample なので適当に insert
	dto.Password = password
	if err := ormgen.Q.User.WithContext(ctx).Create(dto); err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoy) ValidateCredentials(ctx context.Context, email string, password string) (*model.User, error) {
	dto, err := ormgen.User.WithContext(ctx).Where(ormgen.User.Email.Eq(email)).First()
	if err != nil {
		return nil, err
	}
	if dto.Password != password {
		return nil, fmt.Errorf("invalid password")
	}
	return buildDtoToUser(dto), nil
}

func buildUserToDto(user *model.User) *dbschema.User {
	return &dbschema.User{
		ID:        user.ID,
		Name: 		user.Name,
		Email:     user.Email,
	}
}

func buildDtoToUser(dto *dbschema.User) *model.User {
	return &model.User{
		ID:        dto.ID,
		Name: 		dto.Name,
		Email:     dto.Email,
	}
}
