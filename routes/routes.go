package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nico385412/book-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("book/:id", controllers.GetBook)
		v1.GET("book/:id/epub", controllers.GetData)
		v1.GET("book/:id/image", controllers.GetImage)
		v1.GET("books", controllers.GetBooks)
		v1.POST("book", controllers.PostBook)
		v1.DELETE("book/:id", controllers.DeleteBook)
	}

	return r
}
