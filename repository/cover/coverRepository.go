package coverRepository

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nico385412/book-api/config"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func InsertImage(coverData []byte) string {
	fileName := uuid.Must(uuid.NewV4()).String()

	bucket, err := gridfs.NewBucket(
		config.DB.Database("book"),
	)

	uploadStream, err := bucket.OpenUploadStreamWithID(fileName, fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer uploadStream.Close()

	_, err = uploadStream.Write(coverData)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return fileName
}

func DownloadOneImage(id string) bytes.Buffer {
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
