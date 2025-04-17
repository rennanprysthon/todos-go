package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rennanprysthon/go-todo/internal/cases"
	"github.com/rennanprysthon/go-todo/internal/dto"
	"github.com/rennanprysthon/go-todo/internal/web/responses"
)

type CreateTodoHandler struct {
	createTodo cases.CreateTodoCase
}

func NewCreateTodoHandler(createTodoCase cases.CreateTodoCase) *CreateTodoHandler {
	return &CreateTodoHandler{
		createTodo: createTodoCase,
	}
}

func (c *CreateTodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.TodoInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorResponse := responses.ErrorResponse{
			Timestamp: int(time.Now().UnixNano()),
			Status:    500,
			Message:   "Erro interno do sistema",
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	todo := c.createTodo.CreateTodo(input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
