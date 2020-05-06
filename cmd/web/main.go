package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/google"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/memory"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/mongo"

	"github.com/rfinochi/golang-workshop-todo/docs"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title To-Do Sample API
// @version 1.0
// @description Sample To-Do API written in Golang for Go Programming Language Workshop.

// @contact.name Go Programming Language Workshop
// @contact.url https://golang-workshop.io
// @contact.email todoapi@golang-workshop.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/mit-license.php

// @host todo.golang-workshop.io
// @BasePath /api

func main() {
	router := SetupRouter()
	SetupSwagger(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}
	}

	router.Run(fmt.Sprintf(":%s", port))
}

// SetupRouter godoc
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./ui/html", true)))

	api := router.Group("/api")
	api.GET("/", getItemsEndpoint)
	api.GET("/:id", getItemEndpoint)
	api.POST("/", postItemEndpoint)
	api.PUT("/", putItemEndpoint)
	api.PATCH("/:id", updateItemEndpoint)
	api.DELETE("/:id", deleteItemEndpoint)

	return router
}

// SetupSwagger godoc
func SetupSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "./api-docs/index.html")
	})
}

// postItemEndpoint godoc
// @Summary Create a to-do item
// @Description Insert a to-do item into the data store
// @Accept json
// @Produce json
// @Param item body models.Item true "To-Do Item"
// @Success 201 {string} string "{\"message\": \"Ok\"}"
// @Router / [post]
func postItemEndpoint(c *gin.Context) {
	var newItem models.Item
	c.BindJSON(&newItem)

	repo := createRepository()
	repo.CreateItem(newItem)

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
func putItemEndpoint(c *gin.Context) {
	var newItem models.Item
	c.BindJSON(&newItem)

	repo := createRepository()
	repo.CreateItem(newItem)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

// getItemEndpoint godoc
// @Summary Get a to-do item
// @Description Get a to-do item by id from the data store
// @Produce json
// @Param id path int true "To-Do Item Id"
// @Success 200 {object} models.Item
// @Router /{id} [get]
func getItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, repo.GetItem(id))
}

// getItemsEndpoint godoc
// @Summary Get all to-do items
// @Description Get all to-do items from the data store
// @Produce json
// @Success 200 {array} models.Item
// @Router / [get]
func getItemsEndpoint(c *gin.Context) {
	repo := createRepository()
	c.JSON(http.StatusOK, repo.GetItems())
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
func updateItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedItem models.Item
	c.BindJSON(&updatedItem)

	repo := createRepository()
	updatedItem.ID = id
	repo.UpdateItem(updatedItem)

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
func deleteItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	repo.DeleteItem(id)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func createRepository() models.TodoRepository {
	repositoryType := os.Getenv("TODO_REPOSITORY_TYPE")

	if repositoryType == "Mongo" {
		return &mongo.MongoRepository{}
	} else if repositoryType == "Google" {
		return &google.GoogleDatastoreRepository{}
	} else {
		return &memory.MemoryRepository{}
	}
}
