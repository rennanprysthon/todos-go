package cases

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
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

func (c *CreateTodoCase) CreateTodo(todoInput dto.TodoInput) *domain.Todo {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("erro generating uuid: ", err)
	}

	todo := &domain.Todo{
		ID:        uuid.String(),
		Title:     todoInput.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.repository.Save(context.TODO(), todo)

	return todo
}
