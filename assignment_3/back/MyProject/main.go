package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	router.Use(CorsMiddleware())

	router.GET("/items", GetItems)
	router.GET("/items/:id", GetItemById)
	router.POST("/items", AddItem)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	router.Run("localhost:8000")
}
