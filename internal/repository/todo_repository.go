package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rennanprysthon/go-todo/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) *TodoRepository {
	return &TodoRepository{
		collection: collection,
	}
}

func (r *TodoRepository) Save(context context.Context, newTodo *domain.Todo) (*domain.Todo, error) {
	result, err := r.collection.InsertOne(context, newTodo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)

	return newTodo, nil
}

func (r *TodoRepository) FindByUUID(uuid string) *domain.Todo {
	var result domain.Todo
	err := r.collection.FindOne(context.Background(), bson.M{"id": uuid}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		log.Fatal("Erro fetching document", err)
	}

	return &result
}

func (r *TodoRepository) Delete(uuid string) error {
	filter := bson.D{{"id", uuid}}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Update(uuid string, todo *domain.Todo) (*domain.Todo, error) {
	filter := bson.D{{"id", uuid}}

	todo.UpdatedAt = time.Now()

	_, err := r.collection.ReplaceOne(context.TODO(), filter, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
