package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rennanprysthon/go-todo/internal/cases"
	"github.com/rennanprysthon/go-todo/internal/repository"
	"github.com/rennanprysthon/go-todo/internal/web/server"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. ")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	coll := client.Database("todosdb").Collection("todos")

	repository := repository.NewTodoRepository(coll)

	createTodoCase := cases.NewCreateTodoCase(repository)
	addTodoCommentCase := cases.NewAddCommentTodo(repository)
	findTodoByIdCase := cases.NewFindTodoByIdCase(repository)
	deleteTodoCase := cases.NewDeleteTodoCase(repository)

	port := os.Getenv("PORT")
	server := server.NewServer(port)

	server.ConfigureRoutes(
		*createTodoCase,
		*addTodoCommentCase,
		*findTodoByIdCase,
		*deleteTodoCase,
	)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
