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

func logRequestMiddleware(infoLog *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		infoLog.Printf("%s - %s %s %s", c.Request.RemoteAddr, c.Request.Proto, c.Request.Method, c.Request.URL.RequestURI())
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

func tokenAuthMiddleware(errorLog *log.Logger) gin.HandlerFunc {
	requiredToken := os.Getenv(common.APITokenEnvVarName)

	if requiredToken == "" {
		errorLog.Fatal("Please set environment variable ", common.APITokenEnvVarName)
	}

	return func(c *gin.Context) {
		if !strings.Contains(c.Request.RequestURI, "api-docs") {
			token := c.Request.Header.Get(common.APITokenHeaderName)

			if token == "" {
				common.RespondError(c, http.StatusUnauthorized, "API token required")
				return
			}

			if token != requiredToken {
				common.RespondError(c, http.StatusUnauthorized, "Invalid API token (request one via e-mail to rodolfof@shockbyte.software")
				return
			}
		}
		c.Next()
	}
}
