package cases

import (
	"time"

	"github.com/rennanprysthon/go-todo/internal/domain"
	"github.com/rennanprysthon/go-todo/internal/dto"
)

type AddCommentTodoCase struct {
	repository domain.TodoRepository
}

func NewAddCommentTodo(repository domain.TodoRepository) *AddCommentTodoCase {
	return &AddCommentTodoCase{
		repository: repository,
	}
}

func (a *AddCommentTodoCase) AddComment(uuid string, commentInput dto.CommentInput) (*domain.Todo, error) {
	todo := a.repository.FindByUUID(uuid)

	if todo == nil {
		return nil, domain.ErrTodoNotFound
	}

	if todo.Comments == nil {
		todo.Comments = make([]domain.Comment, 0)
	}

	todo.Comments = append(todo.Comments, domain.Comment{
		Username:  commentInput.Username,
		Text:      commentInput.Text,
		CreatedAt: time.Now(),
	})

	todo, err := a.repository.Update(uuid, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
