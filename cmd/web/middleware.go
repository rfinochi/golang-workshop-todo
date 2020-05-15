package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/rfinochi/golang-workshop-todo/pkg/common"
)

func tokenAuthMiddleware(errorLog *log.Logger) gin.HandlerFunc {
	requiredToken := os.Getenv(common.ApiTokenEnvVarName)

	if requiredToken == "" {
		errorLog.Fatal("Please set environment variable ", common.ApiTokenEnvVarName)
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get(common.ApiTokenHeaderName)

		if token == "" {
			common.RespondError(c, http.StatusUnauthorized, "API token required")
			return
		}

		if token != requiredToken {
			common.RespondError(c, http.StatusUnauthorized, "Invalid API token")
			return
		}

		c.Next()
	}
}

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func revisionMiddleware(errorLog *log.Logger) gin.HandlerFunc {
	data, err := ioutil.ReadFile("REVISION")

	if err != nil {
		errorLog.Println("Revision Middleware error: ", err)

		return func(c *gin.Context) {
			c.Next()
		}
	}

	revision := strings.TrimSpace(string(data))

	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Revision", revision)
		c.Next()
	}
}
