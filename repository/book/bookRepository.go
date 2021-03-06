package bookRepository

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/nico385412/book-api/config"
	"github.com/nico385412/book-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func DownloadOneBook(id string) bytes.Buffer {
	fsFiles := config.DB.Database("book").Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	// you can print out the results
	fmt.Println(results)

	bucket, _ := gridfs.NewBucket(
		config.DB.Database("book"),
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(id, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File size to download: %v\n", dStream)
	return buf

}

func GetAllBooks() []*models.Book {
	var books []*models.Book
	cur, err := config.DB.Database("book").Collection("book-info").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
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

func GetBookInfos(id string) models.Book {
	var book models.Book

	err := config.DB.Database("book").Collection("book-info").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)

	if err != nil {
		log.Fatal("Error on Finding document", err)
	}

	return book
}

func InsertBookInfo(book *models.Book) interface{} {
	insertResult, err := config.DB.Database("book").Collection("book-info").InsertOne(context.TODO(), book)
	if err != nil {
		log.Fatal("Error while inserting new Book Infos", err)
	}
	return insertResult.InsertedID
}

func InsertBook(fileName string) int {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	bucket, err := gridfs.NewBucket(
		config.DB.Database("book"),
	)
	uploadStream, err := bucket.OpenUploadStreamWithID(fileName, fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return fileSize
}

func DeleteOneBook(fileName string) {
	config.DB.Database("book").Collection("book-info").DeleteOne(context.TODO(), bson.M{"_id": fileName})

	bucket, err := gridfs.NewBucket(
		config.DB.Database("book"),
	)

	if err != nil {
		log.Fatal(err)
	}

	bucket.Delete(fileName)
}
