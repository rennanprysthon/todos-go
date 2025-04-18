package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rennanprysthon/go-todo/internal/cases"
	"github.com/rennanprysthon/go-todo/internal/web/handlers"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	port   string
}

func NewServer(port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		port:   port,
	}
}

func (s *Server) ConfigureRoutes(
	createTodo cases.CreateTodoCase,
	addCommentTodoCase cases.AddCommentTodoCase,
	findTodoByIdCase cases.FindTodoByIdCase,
	deleteTodoCase cases.DeleteTodoCase,
) {
	createTodoHandler := handlers.NewCreateTodoHandler(createTodo)
	createTodoCommentHandler := handlers.NewAddCommentHandler(addCommentTodoCase)
	newFindTodoByIdHandler := handlers.NewFindTodoByIdHandler(findTodoByIdCase)
	deleteTodoHandler := handlers.NewDeleteTodoHandler(deleteTodoCase)

	s.router.Post("/todos", createTodoHandler.Create)
	s.router.Post("/todos/{todoId}/comments", createTodoCommentHandler.Create)
	s.router.Get("/todos/{todoId}", newFindTodoByIdHandler.FindById)
	s.router.Delete("/todos/{todoId}", deleteTodoHandler.Delete)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
