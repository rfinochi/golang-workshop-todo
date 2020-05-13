package main

import (
	"net/http"
	"strconv"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"

	"github.com/gin-gonic/gin"
)

// getItemEndpoint godoc
// @Summary Get a to-do item
// @Description Get a to-do item by id from the data store
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Success 200 {object} models.Item
// @Router /{id} [get]
func (app *application) getItemEndpoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		app.notFound(c.Writer)
		return
	}

	item, err := app.itemModel.GetItem(id)
	if err == models.ErrNoRecord {
		app.notFound(c.Writer)
		return
	} else if err != nil {
		app.serverError(c.Writer, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, item)
}

// getItemsEndpoint godoc
// @Summary Get all to-do items
// @Description Get all to-do items from the data store
// @Produce json
// @Success 200 {array} models.Item
// @Router / [get]
func (app *application) getItemsEndpoint(c *gin.Context) {
	items, err := app.itemModel.GetItems()
	if err != nil {
		app.serverError(c.Writer, err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

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
	err := c.BindJSON(&newItem)
	if err != nil {
		app.clientError(c.Writer, http.StatusBadRequest)
		return
	}

	err = app.itemModel.CreateItem(newItem)
	if err == models.ErrRecordExist {
		app.conflict(c.Writer, models.ErrRecordExist.Error())
		return
	} else if err != nil {
		app.serverError(c.Writer, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
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
	err := c.BindJSON(&newItem)
	if err != nil {
		app.clientError(c.Writer, http.StatusBadRequest)
		return
	}

	err = app.itemModel.CreateItem(newItem)
	if err == models.ErrRecordExist {
		app.conflict(c.Writer, models.ErrRecordExist.Error())
		return
	} else if err != nil {
		app.serverError(c.Writer, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		app.notFound(c.Writer)
		return
	}

	var updatedItem models.Item
	err = c.BindJSON(&updatedItem)
	if err != nil {
		app.clientError(c.Writer, http.StatusBadRequest)
		return
	}

	updatedItem.ID = id
	err = app.itemModel.UpdateItem(updatedItem)
	if err == models.ErrNoRecord {
		app.notFound(c.Writer)
		return
	} else if err != nil {
		app.serverError(c.Writer, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// deleteItemEndpoint godoc
// @Summary Delete a to-do item
// @Description Delete a to-do item from the data store
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Success 200 {string} string "{\"message\": \"Ok\"}"
// @Router /{id} [delete]
func (app *application) deleteItemEndpoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		app.notFound(c.Writer)
		return
	}

	err = app.itemModel.DeleteItem(id)
	if err != nil {
		app.serverError(c.Writer, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
