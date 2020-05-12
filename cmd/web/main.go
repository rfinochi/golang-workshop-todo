package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/google"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/memory"
	"github.com/rfinochi/golang-workshop-todo/pkg/models/mongo"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	router         *gin.Engine
	itemRepository models.ItemRepository
}

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
	app := &application{
		infoLog:        log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:       log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		itemRepository: createItemRepository(),
	}
	app.initRouter()
	app.addAPIRoutes()
	app.addSwaggerRoutes()

	app.infoLog.Printf("Env PORT '%s'", os.Getenv("PORT"))
	app.infoLog.Printf("Env TODO_REPOSITORY_TYPE '%s'", os.Getenv("TODO_REPOSITORY_TYPE"))
	app.infoLog.Printf("Env TODO_MONGO_URI '%s'", os.Getenv("TODO_MONGO_URI"))

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
		if port == "" {
			port = "8080"
			app.infoLog.Printf("Starting server on port %s", port)
		}
	}

	app.router.Run(fmt.Sprintf(":%s", port))
}

func createItemRepository() models.ItemRepository {
	repositoryType := os.Getenv("TODO_REPOSITORY_TYPE")

	if repositoryType == "Mongo" {
		return &mongo.ItemRepository{}
	} else if repositoryType == "Google" {
		return &google.ItemRepository{}
	} else {
		return &memory.ItemRepository{}
	}
}
