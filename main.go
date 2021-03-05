package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-hata/sample-api/controller"
	"github.com/r-hata/sample-api/middleware"
)

func main() {
	router := gin.Default()
	// Middleware
	router.Use(middleware.RecordUaAndTime)
	// CRUD book
	v1 := router.Group("/v1")
	{
		v1.POST("/books", controller.BookAdd)
		v1.GET("/books", controller.BookList)
		v1.PUT("/books", controller.BookUpdate)
		v1.DELETE("/books", controller.BookDelete)
	}
	router.Run()
}