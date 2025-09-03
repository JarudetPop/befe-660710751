package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student struct
type Book struct {
	Book_ID  string `json:"book_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var books = []Book{
	{Book_ID: "1", Name: "Go Programming", Price: 500, Quantity: 10},
	{Book_ID: "2", Name: "Learning Python", Price: 600, Quantity: 5},
}

func getBooks(c *gin.Context) {
	Book_ID_Query := c.Query("book_id")

	if Book_ID_Query != "" {

		var filtered []Book
		for _, book := range books {
			if fmt.Sprint(book.Book_ID) == Book_ID_Query {
				filtered = append(filtered, book)
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}
	c.JSON(http.StatusOK, books)
}

func main() {
	r := gin.Default()
	r.GET("/books", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "book"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/books", getBooks)
	}

	r.Run(":8080")

}
