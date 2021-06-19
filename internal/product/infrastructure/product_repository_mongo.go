package infrastructure

import (
	"context"
	"grpc-go-templete/internal/product/domain"
	"grpc-go-templete/internal/product/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryMongo struct {
	repository.ProductRepository
	coll *mongo.Collection
}

func NewProductRepositoryMongo(db *mongo.Database) *ProductRepositoryMongo {
	productCollection := db.Collection("products")
	return &ProductRepositoryMongo{
		coll: productCollection,
	}
}

func (p ProductRepositoryMongo) GetOne(ctx context.Context, query interface{}) (data *domain.Product, err error) {
	qResult := p.coll.FindOne(ctx, query)
	data = &domain.Product{}
	err = qResult.Decode(data)
	return
}

func (p ProductRepositoryMongo) GetAll(ctx context.Context, query interface{}, limit, offset int64) (data []*domain.Product, err error) {
	qResult, err := p.coll.Find(ctx, query)
	if err != nil {
		return
	}
	data = make([]*domain.Product, 0)
	err = qResult.Decode(data)
	return
}

func (p ProductRepositoryMongo) Count(ctx context.Context, query interface{}) (data *int64, err error) {
	count, err := p.coll.CountDocuments(ctx, query)
	return &count, err
}

func (p ProductRepositoryMongo) Create(ctx context.Context, entity *domain.Product) (data *domain.Product, err error) {
	insertResult, err := p.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	id := insertResult.InsertedID.(domain.ID)
	qResult := p.coll.FindOne(ctx, &domain.Product{ID: id})
	if err = qResult.Err(); err != nil {
		return
	}
	data = &domain.Product{}
	err = qResult.Decode(data)
	return
}

func (p ProductRepositoryMongo) Update(ctx context.Context, query interface{}, update *domain.Product, limit, offset int64) (data []*domain.Product, err error) {
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
	data = make([]*domain.Product, 0)
	err = qResult.Decode(data)
	return
}
