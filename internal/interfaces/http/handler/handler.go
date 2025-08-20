package handler

import (
	"clean-architecture/internal/interfaces/http/controller"
	"clean-architecture/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

type handler struct {
	controller     *controller.Controller
	authentication *middleware.AuthMiddleware
}

type Handler interface {
	Routes(rg *gin.RouterGroup)
}

type Params struct {
	Controller     *controller.Controller
	Authentication *middleware.AuthMiddleware
}

// @title               Clean Architecture API
// @version             1.0
// @description         Clean Architecture API.
// @servers.url         http://localhost:8080/api/v1
// @servers.description localhost
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewHandler(params *Params) Handler {
	return &handler{
		controller:     params.Controller,
		authentication: params.Authentication,
	}
}

func (h *handler) Routes(rg *gin.RouterGroup) {
	/* register routes */
	h.NewAuthRoutes(rg)
	h.NewUserRoutes(rg)
}
