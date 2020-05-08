package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	router   *gin.Engine
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
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initRouter()
	app.addApiRoutes()
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
