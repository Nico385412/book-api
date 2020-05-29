package main

import (
	"context"
	"log"

	"github.com/nico385412/book-api/config"
	"github.com/nico385412/book-api/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	c := getClient()
	config.DB = c
	err := c.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to database", err)
	} else {
		log.Println("Connected !")
	}

	r := routes.SetupRouter()

	// rc, err := epub.OpenReader("self-dicipline.epub")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rc.Close()
	r.Run(":8000")
}

func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
