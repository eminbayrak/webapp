package db

import (
	"context"
	"go-graphql-mongodb-api/graph/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTodo inserts a new Todo document into MongoDB
func CreateTodo(todo model.NewTodo) error { // Use model.NewTodo type
	// Access the "todos" collection in MongoDB
	todosCollection := client.Database("playground").Collection("todos")

	todoDoc := bson.M{
		"text":   todo.Text,
		"userId": todo.UserID,
	}

	result, err := todosCollection.InsertOne(context.TODO(), todoDoc)
	if err != nil {
		log.Println("Error inserting Todo:", err)
		return err
	}

	log.Println("Inserted new Todo successfully.", result)
	return nil
}

// CreateNews inserts a new News document into MongoDB and returns the result
func CreateNews(news model.NewNews) (*mongo.InsertOneResult, error) {
	// Access the news collection from MongoDB
	newsCollection := client.Database("playground").Collection("news")

	newsDoc := bson.M{
		"title":           news.Title,
		"content":         news.Content,
		"author":          news.Author,
		"category":        news.Category,
		"imageURL":        news.ImageURL,
		"likes":           news.Likes,
		"comments":        news.Comments,
		"share":           news.Share,
		"publicationDate": news.PublicationDate,
	}

	result, err := newsCollection.InsertOne(context.TODO(), newsDoc)
	if err != nil {
		log.Println("Error inserting News:", err)
		return nil, err
	}
	log.Println("Inserted new News successfully.", result)
	return result, nil
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

// FetchNewsByID retrieves a News by its ID from MongoDB.
func FetchNewsByID(id string) (*model.News, error) {
	// Convert the ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ObjectID:", err)
		return nil, err
	}

	// Access the news collection from MongoDB
	newsCollection := client.Database("playground").Collection("news")

	filter := bson.M{"_id": objectID}

	var result model.News

	err = newsCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println("Error fetching News by ID:", err)
		return nil, err
	}

	return &result, nil
}
