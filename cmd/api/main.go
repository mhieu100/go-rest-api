package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "Golang cơ bản", Author: "Gemini"},
    {ID: "2", Title: "Học Microservices", Author: "Google"},
}

func main() {
    router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

    router.GET("/books", func(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, books)
    })

    router.POST("/books", func(c *gin.Context) {
        var newBook Book
        if err := c.BindJSON(&newBook); err != nil {
            return
        }
        books = append(books, newBook)
        c.IndentedJSON(http.StatusCreated, newBook)
    })

    router.Run("localhost:8080")
}