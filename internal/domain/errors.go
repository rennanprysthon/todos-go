package domain

import "errors"

var (
	ErrTodoNotFound = errors.New("Todo nao encontrado")
	ErrDeleteTodo   = errors.New("Erro ao deletar o todo")
	ErrCreateTodo   = errors.New("Erro ao criar todo")
)
