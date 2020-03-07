package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"mygraphql/graph/generated"
	"mygraphql/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, limit *int) ([]*model.User, error) {
	findOptions := options.Find()
	if limit != nil {
		limit64 := int64(*limit)
		findOptions.SetLimit(limit64)
	}

	var userArr []*model.User

	cursor, err := NewMongoClient.Database("Account").Collection("users").Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		userArr = append(userArr, &user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	return userArr, nil
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.Hex(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) Pass(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
