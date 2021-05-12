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
	userCollection := db.Collection("user")
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
	err = qResult.Decode(data)
	return
}

func (p userPersistance) Count(query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(context.Background(), query)
	return &count, err
}

func (p userPersistance) Save(entity *model.User) (data *model.User, err error) {
	id := entity.ID
	if id == "" {
		qResult, err := p.coll.InsertOne(context.Background(), entity)
		if err != nil {
			return nil, err
		}
		id = qResult.InsertedID.(string)
	} else {
		_, err := p.coll.UpdateByID(context.Background(), entity.ID, entity)
		if err != nil {
			return nil, err
		}
	}
	// qResult, err := p.coll.FindOne(&model.User{ID: id})
}

func (p userPersistance) Update(query interface{}, update *model.User, options interface{}) (data []*model.User, err error) {
	return []*model.User{}, nil
}