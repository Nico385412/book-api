package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nico385412/book-api/converter"
	"github.com/nico385412/book-api/repository"
	uuid "github.com/satori/go.uuid"
)

func GetBooks(c *gin.Context) {
	books := repository.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	fileUUID := c.Param("id")
	data := repository.DownloadOneBook(fileUUID)
	c.Data(http.StatusOK, "application/epub+zip", data.Bytes())
}

func DeleteBook(c *gin.Context) {
	bookUUID := c.Param("id")
	repository.DeleteOneBook(bookUUID)
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

	repository.InsertBook(fileName)

	book := converter.ConvertFileToBookModel(&fileName)
	repository.InsertBookInfo(book)

	c.JSON(http.StatusOK, gin.H{
		"inserted": fileName,
	})
}
