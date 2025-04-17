package cases

import "github.com/rennanprysthon/go-todo/internal/domain"

type DeleteTodoCase struct {
	repository domain.TodoRepository
}

func NewDeleteTodoCase(repository domain.TodoRepository) *DeleteTodoCase {
	return &DeleteTodoCase{
		repository: repository,
	}
}

func (d *DeleteTodoCase) Delete(uuid string) error {
	todo := d.repository.FindByUUID(uuid)

	if todo == nil {
		return domain.ErrTodoNotFound
	}

	err := d.repository.Delete(uuid)
	if err != nil {
		return domain.ErrDeleteTodo
	}

	return nil
}
