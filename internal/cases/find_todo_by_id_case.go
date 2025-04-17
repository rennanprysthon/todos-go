package cases

import "github.com/rennanprysthon/go-todo/internal/domain"

type FindTodoByIdCase struct {
	repository domain.TodoRepository
}

func NewFindTodoByIdCase(todoRepository domain.TodoRepository) *FindTodoByIdCase {
	return &FindTodoByIdCase{
		repository: todoRepository,
	}
}

func (f *FindTodoByIdCase) FindByID(uuid string) (*domain.Todo, error) {
	todo := f.repository.FindByUUID(uuid)

	if todo == nil {
		return nil, domain.ErrTodoNotFound
	}

	return todo, nil
}
