package domain

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Username  string
	Text      string
	CreatedAt time.Time
}

type Todo struct {
	ID          string
	Title       string
	Comments    []Comment
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
}

func (t *Todo) AddComment(username, text string) {
	t.Comments = append(t.Comments, Comment{
		Username:  username,
		Text:      text,
		CreatedAt: time.Now(),
	})

	t.UpdatedAt = time.Now()
}

func NewTodo(title string) *Todo {
	todo := &Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Comments:  make([]Comment, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return todo
}
