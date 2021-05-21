package persistance

import (
	"context"
	"grpc-go-templete/internal/user/model"
	"grpc-go-templete/internal/user/repository"

	"go.mongodb.org/mongo-driver/mongo"
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

func (p userPersistance) GetOne(query interface{}) (data *model.User, err error) {
	qResult := p.coll.FindOne(context.Background(), query)
	err = qResult.Decode(data)
	return
}

func (p userPersistance) GetAll(query, options interface{}) (data []*model.User, err error) {
	qResult, err := p.coll.Find(context.Background(), query)
	if err != nil {
		return
	}
	data = make([]*model.User, 0)
	err = qResult.Decode(data)
	return
}

func (p userPersistance) Count(query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(context.Background(), query)
	return &count, err
}

func (p userPersistance) Create(entity *model.User) (data *model.User, err error) {
	insertResult, err := p.coll.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, err
	}
	id := insertResult.InsertedID.(model.ID)
	qResult := p.coll.FindOne(context.Background(), &model.User{ID: id})
	if err = qResult.Err(); err != nil {
		return
	}
	data = &model.User{}
	err = qResult.Decode(data)
	return
}

func (p userPersistance) Update(query interface{}, update *model.User, options interface{}) (data []*model.User, err error) {
	return []*model.User{}, nil
}
