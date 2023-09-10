package db

import (
	"context"
	"go-graphql-mongodb-api/graph/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateTodo inserts a new Todo document into MongoDB
func CreateTodo(todo model.NewTodo) error { // Use model.NewTodo type
	// Access the "todos" collection in MongoDB
	todosCollection := client.Database("playground").Collection("todos")

	todoDoc := bson.M{
		"text":   todo.Text,
		"userId": todo.UserID,
	}

	_, err := todosCollection.InsertOne(context.TODO(), todoDoc)
	if err != nil {
		log.Println("Error inserting Todo:", err)
		return err
	}

	log.Println("Inserted new Todo successfully")
	return nil
}

// FetchTodoByText retrieves a Todo by its text field from MongoDB.
func FetchTodoByText(text string) (*model.Todo, error) {
	todosCollection := client.Database("playground").Collection("todos")

	filter := bson.M{"text": text}

	var result model.Todo

	err := todosCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
