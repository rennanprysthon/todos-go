package domain

import "errors"

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrDeleteTodo   = errors.New("Erro ao deletar o todo")
)
