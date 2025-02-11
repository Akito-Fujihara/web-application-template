package repository

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	domainRepo "github.com/Akito-Fujihara/web-application-template/app/domain/repository"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/dbschema"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/ormgen"
)

type TodoRepository struct {}

func NewTodoRepository() domainRepo.ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	dto := buildTodoToDto(todo)
	if err := ormgen.Q.Todo.WithContext(ctx).Create(dto); err != nil {
		return err
	}
	return nil
}

func buildTodoToDto(todo *model.Todo) *dbschema.Todo {
	return &dbschema.Todo{
		ID:        todo.ID,
		UserID: 	todo.UserID,
		Title:     todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
	}
}
