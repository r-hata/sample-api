package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-hata/sample-api/controller"
	"github.com/r-hata/sample-api/logger"
	"github.com/r-hata/sample-api/middleware"
)

var port string

func main() {
	defer logger.Sync()
	router := gin.Default()
	// Middleware
	router.Use(middleware.RecordUaAndTime)
	// CRUD book
	v1 := router.Group("/v1")
	{
		v1.POST("/books", controller.BookAdd)
		v1.GET("/books/:id", controller.BookGet)
		v1.GET("/books", controller.BookList)
		v1.PUT("/books", controller.BookUpdate)
		v1.DELETE("/books/:id", controller.BookDelete)
	}
	router.Run(port)
}

func init() {
	mode := os.Getenv("GIN_MODE")
	port = ":8080"
	if mode == "release" {
		port = ":80"
	}
}
