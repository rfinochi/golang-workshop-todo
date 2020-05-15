package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rfinochi/golang-workshop-todo/pkg/common"
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
	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()
	app.addSwaggerRoutes()

	app.infoLog.Printf("Env %s '%s'", common.PortEnvVarName, os.Getenv(common.PortEnvVarName))
	app.infoLog.Printf("Env %s '%s'", common.PortEnvVarName2, os.Getenv(common.PortEnvVarName2))
	app.infoLog.Printf("Env %s '%s'", common.ApiTokenEnvVarName, os.Getenv(common.ApiTokenEnvVarName))
	app.infoLog.Printf("Env %s '%s'", common.RepositoryMongoURIEnvVarName, os.Getenv(common.RepositoryMongoURIEnvVarName))
	app.infoLog.Printf("Env %s '%s'", common.ApiTokenEnvVarName, os.Getenv(common.ApiTokenEnvVarName))

	port := os.Getenv(common.PortEnvVarName)
	if port == "" {
		port = os.Getenv(common.PortEnvVarName2)
		if port == "" {
			port = common.PortDefault
			app.infoLog.Printf("Starting server on port %s", port)
		}
	}

	app.router.Run(fmt.Sprintf(":%s", port))
}
