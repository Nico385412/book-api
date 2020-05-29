package usecase

import (
	"context"
	"github.com/nico385412/book-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BookUseCase struct {
	db *mongo.Database
}

func (b *BookUseCase) New(mongo *mongo.Client) {
	b.db = mongo.Database("books")
}

func (b *BookUseCase) ReturnAllBooks(filter bson.M) []*models.Book {
	var books []*models.Book
	collection := b.db.Collection("book-info")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all documents", err)
	}
	for cur.Next(context.TODO()) {
		var book models.Book
		err = cur.Decode(&book)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		books = append(books, &book)
	}
	return books
}
