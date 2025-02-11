package repository

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
)
type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User, password string) error
}

type ITodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
}
