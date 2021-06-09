package infrastructure

import (
	"context"
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserPersistance is persistance model
type UserPersistance struct {
	repository.UserRepository
	coll *mongo.Collection
}

// NewUserPersistance return new UserPersistance instance
func NewUserPersistance(db *mongo.Database) *UserPersistance {
	userCollection := db.Collection("users")
	return &UserPersistance{
		coll: userCollection,
	}
}

// CreateIndexes method to create DB indexes
func (p UserPersistance) CreateIndexes() {
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{primitive.E{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	p.coll.Indexes().CreateMany(context.TODO(), indexModels)
}

// GetOne is method to get one User
func (p UserPersistance) GetOne(ctx context.Context, query interface{}) (data *domain.User, err error) {
	qResult := p.coll.FindOne(ctx, query)
	data = &domain.User{}
	err = qResult.Decode(data)
	return
}

//GetAll is method to get Users
func (p UserPersistance) GetAll(ctx context.Context, query interface{}, limit, offset int64) (data []*domain.User, err error) {
	qResult, err := p.coll.Find(ctx, query)
	if err != nil {
		return
	}
	data = make([]*domain.User, 0)
	err = qResult.Decode(data)
	return
}

// Count is method to count Users
func (p UserPersistance) Count(ctx context.Context, query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(ctx, query)
	return &count, err
}

// Create is method to create Users
func (p UserPersistance) Create(ctx context.Context, entity *domain.User) (data *domain.User, err error) {
	insertResult, err := p.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	id := insertResult.InsertedID.(domain.ID)
	qResult := p.coll.FindOne(ctx, &domain.User{ID: id})
	if err = qResult.Err(); err != nil {
		return
	}
	data = &domain.User{}
	err = qResult.Decode(data)
	return
}

// Update is method to create User(s)
func (p UserPersistance) Update(ctx context.Context, query interface{}, update *domain.User, limit, offset int64) (data []*domain.User, err error) {
	updateResult, err := p.coll.UpdateMany(ctx, query, update)
	if err != nil {
		return
	}
	if updateResult.UpsertedCount == 0 {
		return
	}
	qResult, err := p.coll.Find(ctx, query)
	if err != nil {
		return
	}
	data = make([]*domain.User, 0)
	err = qResult.Decode(data)
	return
}
