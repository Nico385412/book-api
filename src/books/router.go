package books

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nico385412/book-api/books/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPing should return ping string
func GetPing(c *context.Context, mongo *mongo.Client) {
	bookService := usecase.BookUseCase{}
	bookService.New(mongo)
	c.JSON(200, gin.H{
		"serverTime": bookService.ReturnAllBooks(bson.M{}),
	})
}
