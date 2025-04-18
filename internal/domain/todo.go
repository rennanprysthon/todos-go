package domain

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Username  string
	Text      string
	CreatedAt time.Time
}

type Todo struct {
	mu          sync.RWMutex
	ID          string
	Title       string
	Comments    []Comment
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
}

func (t *Todo) AddComment(username, text string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Comments != nil {
		t.Comments = make([]Comment, 0)
	}

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
