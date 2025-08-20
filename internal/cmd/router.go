package cmd

import (
	"clean-architecture/internal/interfaces/http/handler"
	"clean-architecture/internal/interfaces/http/middleware"

	_ "clean-architecture/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *app) newRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.NewGinMiddleware())

	/* swagger */
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	/* health check */
	v1.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	router := handler.NewHandler(&handler.Params{
		Controller:     a.controller,
		Authentication: a.authentication,
	})
	router.Routes(v1)

	return r
}
