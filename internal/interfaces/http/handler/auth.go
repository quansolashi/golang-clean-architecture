package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) NewAuthRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/auth")
	r.POST("/login", h.controller.Login)
}
