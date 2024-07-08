package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"errors"
)

type book struct {
	// Conversion for JSON format of the books informations.
	ID 	 string `json:"id"`
	Title 	 string `json:"title"`
	Author 	 string `json:"author"`
	Quantity int    `json:"quantity"`
}

func main() {
	shelf := []book{
		{ID: "1", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
		{ID: "2", Title: "1984", Author: "George Orwell", Quantity: 10},
		{ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 7},
		{ID: "4", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 4},
		{ID: "5", Title: "Moby-Dick", Author: "Herman Melville", Quantity: 6},
		{ID: "6", Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 8},
		{ID: "7", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
		{ID: "8", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 9},
		{ID: "9", Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Quantity: 2},
		{ID: "10", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 1},
	}

	// gin.Context: informations about the request from the client.
	// This is a closure for using the variable from the main func.
	getBooks := func(c *gin.Context) {
		// Returning a OK with the information from the slice shelf.
		c.IndentedJSON(http.StatusOK, shelf)
	}

	createBook := func(c *gin.Context) {
		var newBook book

		if err := c.BindJSON(&newBook); err != nil {
			return
		}

		shelf = append(shelf, newBook)
		c.IndentedJSON(http.StatusCreated, newBook)
	}

	router := gin.Default()
	// Handling route localhost:8080/books
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
	
}
