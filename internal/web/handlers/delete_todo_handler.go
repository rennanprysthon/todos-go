package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rennanprysthon/go-todo/internal/cases"
	"github.com/rennanprysthon/go-todo/internal/domain"
	"github.com/rennanprysthon/go-todo/internal/web/responses"
)

type DeleteTodoHandler struct {
	deleteTodoCase cases.DeleteTodoCase
}

func NewDeleteTodoHandler(deleteTodoCase cases.DeleteTodoCase) *DeleteTodoHandler {
	return &DeleteTodoHandler{
		deleteTodoCase: deleteTodoCase,
	}
}

func (d *DeleteTodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")

	err := d.deleteTodoCase.Delete(todoId)

	if err == domain.ErrTodoNotFound {
		errorResponse := responses.ErrorResponse{
			Timestamp: int(time.Now().UnixNano()),
			Status:    404,
			Message:   "Objeto com id " + todoId + " nao foi encontrado",
		}
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(errorResponse)

		return
	}

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

	w.WriteHeader(http.StatusNoContent)
}
