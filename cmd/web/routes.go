package main

import (
	"net/http"

	"github.com/rfinochi/golang-workshop-todo/docs"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./ui/html", true)))

	api := router.Group("/api")
	api.GET("/", app.getItemsEndpoint)
	api.GET("/:id", app.getItemEndpoint)
	api.POST("/", app.postItemEndpoint)
	api.PUT("/", app.putItemEndpoint)
	api.PATCH("/:id", app.updateItemEndpoint)
	api.DELETE("/:id", app.deleteItemEndpoint)

	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "./api-docs/index.html")
	})

	return router
}
