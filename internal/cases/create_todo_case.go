package cases

import (
	"context"

	"github.com/rennanprysthon/go-todo/internal/domain"
	"github.com/rennanprysthon/go-todo/internal/dto"
)

type CreateTodoCase struct {
	repository domain.TodoRepository
}

func NewCreateTodoCase(todoRepository domain.TodoRepository) *CreateTodoCase {
	return &CreateTodoCase{
		repository: todoRepository,
	}
}

func (c *CreateTodoCase) CreateTodo(todoInput dto.TodoInput) (*domain.Todo, error) {
	todo := domain.NewTodo(todoInput.Title)

	newTodo, err := c.repository.Save(context.Background(), todo)
	if err != nil {
		return nil, domain.ErrCreateTodo
	}

	return newTodo, nil
}
