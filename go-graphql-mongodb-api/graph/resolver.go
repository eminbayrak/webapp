package graph

import "go-graphql-mongodb-api/graph/model"

type Resolver struct {
	// Use something from the model package
	UserModel model.User
}
