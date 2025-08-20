package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) NewUserRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.Use(h.authentication.Authenticate())
	r.GET("", h.controller.ListUsers)
	r.GET("/:userId", h.controller.GetUser)
}
