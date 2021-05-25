package persistance

import (
	"grpc-go-templete/internal/auth/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type accountPersistance struct {
	repository.AccountRepository
	coll *mongo.Collection
}

func NewAccountPersistance(db *mongo.Database) *accountPersistance {
	accountCollection := db.Collection("accounts")
	return &accountPersistance{
		coll: accountCollection,
	}
}
