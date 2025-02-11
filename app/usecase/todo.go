package usecase

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/Akito-Fujihara/web-application-template/app/domain/repository"
)

type ITodoUsecase interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
}

type TodoUsecase struct {
	todoRepo repository.ITodoRepository
}

func NewTodoUsecase(todoRepo repository.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{
		todoRepo: todoRepo,
	}
}

func (u *TodoUsecase) CreateTodo(ctx context.Context, todo *model.Todo) error {
	return u.todoRepo.CreateTodo(ctx, todo)
}
