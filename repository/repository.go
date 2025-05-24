package repository

import (
	"context"
	"go-product-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(id string) (*models.Product, error)
	Create(product models.Product) (*models.Product, error)
	Update(id string, product models.Product) error
	Delete(id string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(col *mongo.Collection) ProductRepository {
	return &productRepository{collection: col}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var p models.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) GetById(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product models.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	return &product, err
}

func (r *productRepository) Create(product models.Product) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.Id = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, product)
	return &product, err
}

func (r *productRepository) Update(id string, product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"price": product.Price,
		},
	}
	_, err = r.collection.UpdateByID(ctx, objID, update)
	return err
}

func (r *productRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
