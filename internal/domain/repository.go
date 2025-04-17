package domain

import "context"

type TodoRepository interface {
	Save(context context.Context, todo *Todo) (*Todo, error)
	FindByUUID(uuid string) *Todo
	Delete(uuid string) error
	Update(uuid string, todo *Todo) (*Todo, error)
}
