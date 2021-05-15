package persistance

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProvider struct {
	client *mongo.Client
}

func NewMongoProvider(uri string) (provider *MongoProvider, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	provider = &MongoProvider{}
	provider.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	return
}

func (p MongoProvider) GetDatabase(name string) *mongo.Database {
	return p.client.Database(name)
}
