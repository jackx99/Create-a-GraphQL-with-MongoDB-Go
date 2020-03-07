package graph

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var NewMongoClient *mongo.Client

func NewDatabase()  {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://john:1234@cluster0-vhryk.gcp.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	NewMongoClient = client
}