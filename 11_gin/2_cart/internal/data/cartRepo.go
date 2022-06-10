package data

import (
	"context"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cartRepo struct {
	db *mongo.Database
}

func NewCartRepo(db *mongo.Database) CartRepo {
	return &cartRepo{db: db}
}

func (c *cartRepo) UpsertCart(ctx context.Context, cart models.Cart) models.Cart {
	if _, err := c.cart().ReplaceOne(
		ctx,
		bson.M{"id": cart.Id},
		cart,
		options.Replace().SetUpsert(true),
	); err != nil {
		panic(err)
	}
	return cart
}

func (c *cartRepo) cart() *mongo.Collection {
	return c.db.Collection("cart")
}
