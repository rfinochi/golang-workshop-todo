package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondOk godoc
func RespondOk(c *gin.Context) {
	RespondOkWithData(c, gin.H{"message": "OK"})
}

// RespondOkWithData godoc
func RespondOkWithData(c *gin.Context, data interface{}) {
	RespondWithData(c, http.StatusOK, data)
}

// RespondCreated godoc
func RespondCreated(c *gin.Context) {
	RespondWithData(c, http.StatusCreated, gin.H{"message": "OK"})
}

// RespondWithData godoc
func RespondWithData(c *gin.Context, httpStatusCode int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(httpStatusCode, data)
}

// RespondError godoc
func RespondError(c *gin.Context, httpStatusCode int, message interface{}) {
	c.Header("Content-Type", "application/json")
	c.AbortWithStatusJSON(httpStatusCode, gin.H{"error": message})
}
