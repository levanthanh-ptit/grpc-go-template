package persistance

import (
	"context"
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userPersistance struct {
	repository.UserRepository
	coll *mongo.Collection
}

func NewUserPersistance(db *mongo.Database) *userPersistance {
	userCollection := db.Collection("users")
	return &userPersistance{
		coll: userCollection,
	}
}

func (p userPersistance) CreateIndexes() {
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{primitive.E{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	p.coll.Indexes().CreateMany(context.TODO(), indexModels)
}

func (p userPersistance) GetOne(query interface{}) (data *domain.User, err error) {
	qResult := p.coll.FindOne(context.Background(), query)
	err = qResult.Decode(data)
	return
}

func (p userPersistance) GetAll(query interface{}) (data []*domain.User, err error) {
	qResult, err := p.coll.Find(context.Background(), query)
	if err != nil {
		return
	}
	data = make([]*domain.User, 0)
	err = qResult.Decode(data)
	return
}

func (p userPersistance) Count(query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(context.Background(), query)
	return &count, err
}

func (p userPersistance) Create(entity *domain.User) (data *domain.User, err error) {
	insertResult, err := p.coll.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, err
	}
	id := insertResult.InsertedID.(domain.ID)
	qResult := p.coll.FindOne(context.Background(), &domain.User{ID: id})
	if err = qResult.Err(); err != nil {
		return
	}
	data = &domain.User{}
	err = qResult.Decode(data)
	return
}

func (p userPersistance) Update(query interface{}, update *domain.User) (data []*domain.User, err error) {
	updateResult, err := p.coll.UpdateMany(context.Background(), query, update)
	if err != nil {
		return
	}
	if updateResult.UpsertedCount == 0 {
		return
	}
	qResult, err := p.coll.Find(context.Background(), query)
	if err != nil {
		return
	}
	data = make([]*domain.User, 0)
	err = qResult.Decode(data)
	return
}
