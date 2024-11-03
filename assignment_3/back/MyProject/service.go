package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")

	// I check if Item exists
	for _, item := range items {

		// If Item exists, it will be returned
		if id == fmt.Sprintf("%d", item.Id) {
			c.IndentedJSON(http.StatusOK, item)

			// If Item is not founded, the function will be stopped after the server responded
			return
		}
	}

	// If function was not stopped, it means that Item was not found. The corresponding message will be responded by the server
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item with id " + id + " was not found"})
}

func AddItem(c *gin.Context) {
	var item Item
	err := c.BindJSON(&item)

	// If error with parsing JSON occurs, the Item will not be added
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item.Id = len(items) + 1
	items = append(items, item)
	c.IndentedJSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	// Before I update Item's data, I need to check if my item exists
	id := c.Param("id")

	for index, item := range items {

		// if I find the Item by Id, I can update its data
		if id == fmt.Sprintf("%d", item.Id) {
			var item Item

			err := c.BindJSON(&item)

			// If error with parsing JSON occurs, the Item will not be added
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}

			items[index] = item

			c.IndentedJSON(http.StatusOK, item)

			// After the procedure of updating will be finished, the function execution will be stopped
			return
		}
	}

	// If function was not stopped, it means that Item was not found. The corresponding message will be responded by the server
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item  with id " + id + " was not found"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	//
	//fmt.Println(id)

	// I check if Item exists
	for index, item := range items {

		// If Item exists, it can be deleted
		if id == fmt.Sprintf("%d", item.Id) {

			// I delete element from array using slices
			items = append(items[:index], items[index+1:]...)

			// After server response with the message, that Item has successfully been deleted
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Item " + item.Name + " has been deleted"})

			// After function finishes (stops)
			return
		}
	}

	// If function was not stopped, it means that Item was not found. The corresponding message will be responded by the server
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item with id " + id + " was not found"})
}
