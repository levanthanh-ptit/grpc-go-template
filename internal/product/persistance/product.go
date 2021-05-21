package persistance

import (
	"context"
	"grpc-go-templete/internal/product/model"
	"grpc-go-templete/internal/product/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type productPersistance struct {
	repository.ProductRepository
	coll *mongo.Collection
}

func NewProductPersistance(db *mongo.Database) *productPersistance {
	productCollection := db.Collection("products")
	return &productPersistance{
		coll: productCollection,
	}
}

func (p productPersistance) GetOne(query interface{}) (data *model.Product, err error) {
	qResult := p.coll.FindOne(context.Background(), query)
	err = qResult.Decode(data)
	return
}

func (p productPersistance) GetAll(query, options interface{}) (data []*model.Product, err error) {
	qResult, err := p.coll.Find(context.Background(), query)
	if err != nil {
		return
	}
	data = make([]*model.Product, 0)
	err = qResult.Decode(data)
	return
}

func (p productPersistance) Count(query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(context.Background(), query)
	return &count, err
}

func (p productPersistance) Create(entity *model.Product) (data *model.Product, err error) {
	insertResult, err := p.coll.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, err
	}
	id := insertResult.InsertedID.(model.ID)
	qResult := p.coll.FindOne(context.Background(), &model.Product{ID: id})
	if err = qResult.Err(); err != nil {
		return
	}
	data = &model.Product{}
	err = qResult.Decode(data)
	return
}

func (p productPersistance) Update(query interface{}, update *model.Product, options interface{}) (data []*model.Product, err error) {
	return []*model.Product{}, nil
}
