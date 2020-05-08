package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/google"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/memory"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/mongo"

	"github.com/gin-gonic/gin"
)

// postItemEndpoint godoc
// @Summary Create a to-do item
// @Description Insert a to-do item into the data store
// @Accept json
// @Produce json
// @Param item body models.Item true "To-Do Item"
// @Success 201 {string} string "{\"message\": \"Ok\"}"
// @Router / [post]
func (app *application) postItemEndpoint(c *gin.Context) {
	var newItem models.Item
	c.BindJSON(&newItem)

	repo := createRepository()
	err := repo.CreateItem(newItem)
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusCreated, gin.H{"message": "OK"})
	}
}

// putItemEndpoint godoc
// @Summary Create a to-do item
// @Description Insert a to-do item into the data store
// @Accept json
// @Produce json
// @Param item body models.Item true "To-Do Item"
// @Success 201 {string} string "{\"message\": \"Ok\"}"
// @Router / [put]
func (app *application) putItemEndpoint(c *gin.Context) {
	var newItem models.Item
	c.BindJSON(&newItem)

	repo := createRepository()
	err := repo.CreateItem(newItem)
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusCreated, gin.H{"message": "OK"})
	}
}

// getItemEndpoint godoc
// @Summary Get a to-do item
// @Description Get a to-do item by id from the data store
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Success 200 {object} models.Item
// @Router /{id} [get]
func (app *application) getItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	item, err := repo.GetItem(id)
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, item)
	}
}

// getItemsEndpoint godoc
// @Summary Get all to-do items
// @Description Get all to-do items from the data store
// @Produce json
// @Success 200 {array} models.Item
// @Router / [get]
func (app *application) getItemsEndpoint(c *gin.Context) {
	repo := createRepository()
	items, err := repo.GetItems()
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

// updateItemEndpoint godoc
// @Summary Update a to-do item
// @Description Update a to-do item into the data store
// @Accept json
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Param item body models.Item true "To-Do Item"
// @Success 200 {string} string "{\"message\": \"Ok\"}"
// @Router /{id} [patch]
func (app *application) updateItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedItem models.Item
	c.BindJSON(&updatedItem)

	repo := createRepository()
	updatedItem.ID = id
	err := repo.UpdateItem(updatedItem)
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

// deleteItemEndpoint godoc
// @Summary Delete a to-do item
// @Description Delete a to-do item from the data store
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Success 200 {string} string "{\"message\": \"Ok\"}"
// @Router /{id} [delete]
func (app *application) deleteItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	err := repo.DeleteItem(id)
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func createRepository() models.TodoRepository {
	repositoryType := os.Getenv("TODO_REPOSITORY_TYPE")

	if repositoryType == "Mongo" {
		return &mongo.MongoRepository{}
	} else if repositoryType == "Google" {
		return &google.GoogleRepository{}
	} else {
		return &memory.MemoryRepository{}
	}
}
