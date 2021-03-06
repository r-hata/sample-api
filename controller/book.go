package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/r-hata/sample-api/logger"
	"github.com/r-hata/sample-api/model"
	"github.com/r-hata/sample-api/service"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bookService := service.BookService{}
	err = bookService.Add(&book)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   book,
	})
}

func BookGet(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bookService := service.BookService{}
	has, book, err := bookService.Get(int(intID))
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if !has {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   book,
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	list := bookService.List()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    list,
	})
}

func BookUpdate(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bookService := service.BookService{}
	err = bookService.Update(&book)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.Status(http.StatusNoContent)
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bookService := service.BookService{}
	err = bookService.Delete(int(intID))
	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.Status(http.StatusNoContent)
}
