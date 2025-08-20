package controller

import (
	"clean-architecture/internal/domain/dto"
	"clean-architecture/internal/interfaces/http/response"
	"clean-architecture/internal/interfaces/http/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     login
// @Description login
// @Tags        Auth
// @Router      /auth/login [post]
// @Param       loginRequest body dto.LoginRequest true "login request"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.AuthResponse
// @Failure     400 {object} util.ErrorResponse
// @Failure     401 {object} util.ErrorResponse
// @Failure     500 {object} util.ErrorResponse
func (c *Controller) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.badRequest(ctx, err)
		return
	}

	user, token, err := c.Usecase.Auth.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		c.unauthorized(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth: service.NewAuth(user, token).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}
