package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	bookRepository "github.com/nico385412/book-api/repository/book"
	coverRepository "github.com/nico385412/book-api/repository/cover"
	"github.com/nico385412/book-api/service"
	uuid "github.com/satori/go.uuid"
)

func GetBooks(c *gin.Context) {
	books := bookRepository.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	fileUUID := c.Param("id")
	data := bookRepository.GetBookInfos(fileUUID)
	c.JSON(http.StatusOK, data)
}

func GetData(c *gin.Context) {
	fileUUID := c.Param("id")
	data := coverRepository.DownloadOneImage(fileUUID)
	c.Data(http.StatusOK, "application/epub+zip", data.Bytes())
}

func GetImage(c *gin.Context) {
	fileUUID := c.Param("id")
	book := bookRepository.GetBookInfos(fileUUID)
	data := coverRepository.DownloadOneImage(book.CoverID)
	c.Data(http.StatusOK, "", data.Bytes())
}

func DeleteBook(c *gin.Context) {
	bookUUID := c.Param("id")
	bookRepository.DeleteOneBook(bookUUID)
}

func PostBook(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint("get form err : %s", err.Error()))
		return
	}

	fileName := uuid.Must(uuid.NewV4()).String()
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err : %s", err.Error()))
		return
	}

	//Insert Ebook File
	bookRepository.InsertBook(fileName)

	bookImage, err2 := service.GetCover(&fileName)
	var coverId string

	if err2 != nil {
		log.Printf("There is no image for this cover")
	} else {
		coverId = coverRepository.InsertImage(bookImage)
	}

	book := service.ConvertFileToBookModel(&fileName, &coverId)
	bookRepository.InsertBookInfo(book)

	c.JSON(http.StatusOK, gin.H{
		"inserted": fileName,
	})

	defer os.Remove(fileName)
}
