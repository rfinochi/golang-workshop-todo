package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rfinochi/golang-workshop-todo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) initRouter() {
	app.router = gin.Default()
}

func (app *application) addAPIRoutes() {
	if app.router != nil {
		app.router.Use(logRequestMiddleware(app.infoLog))
		app.router.Use(revisionMiddleware(app.errorLog))
		app.router.Use(requestIDMiddleware())
		app.router.Use(tokenAuthMiddleware(app.errorLog))
		app.router.Use(static.Serve("/", static.LocalFile("./ui/html", true)))

		api := app.router.Group("/api")
		api.GET("/", app.getItemsEndpoint)
		api.GET("/:id", app.getItemEndpoint)
		api.POST("/", app.postItemEndpoint)
		api.PUT("/", app.putItemEndpoint)
		api.PATCH("/:id", app.updateItemEndpoint)
		api.DELETE("/:id", app.deleteItemEndpoint)
	}
}

func (app *application) addSwaggerRoutes() {
	if app.router != nil {
		docs.SwaggerInfo_swagger.Schemes = []string{"https"}
		app.router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		app.router.GET("/api-docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "./api-docs/index.html")
		})
	}
}
