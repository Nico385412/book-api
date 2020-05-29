package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Client

func books() *mongo.Collection {
	return DB.Database("book").Collection("book-infos")
}
