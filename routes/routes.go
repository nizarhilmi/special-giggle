package routes

import (
	"api-go-basic/controller"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/movies", controller.GetAllMovies)
	r.POST("/movies", controller.CreateMovie)

	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	return r
}
