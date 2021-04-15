// controllers/books.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oobohP/goback/models"
)

// FindBooks fetches all books from the database
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBooks creates a book in the database
func CreateBook(c *gin.Context) {
	// Validation
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil { // ShouldBindJSON checks the request body and passes the schema *our model with bindings*
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
