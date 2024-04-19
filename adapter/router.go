/*
Package adapter adapts to all kinds of framework or protocols.
*/
package adapter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"literank.com/fulltext-books/application"
	"literank.com/fulltext-books/application/executor"
	"literank.com/fulltext-books/domain/model"
)

const fieldQuery = "q"

// RestHandler handles all restful requests
type RestHandler struct {
	bookOperator *executor.BookOperator
}

func newRestHandler(wireHelper *application.WireHelper) *RestHandler {
	return &RestHandler{
		bookOperator: executor.NewBookOperator(wireHelper.BookManager()),
	}
}

// MakeRouter makes the main router
func MakeRouter(wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := newRestHandler(wireHelper)
	// Create a new Gin router
	r := gin.Default()

	r.GET("/books", rest.searchBooks)
	r.POST("/books", rest.createBook)
	return r, nil
}

// Create a new book
func (r *RestHandler) createBook(c *gin.Context) {
	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, err := r.bookOperator.CreateBook(c, &reqBody)
	if err != nil {
		fmt.Printf("Failed to create: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": bookID})
}

func (r *RestHandler) searchBooks(c *gin.Context) {
	books, err := r.bookOperator.SearchBooks(c, c.Query(fieldQuery))
	if err != nil {
		fmt.Printf("Failed to search books: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to search books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
