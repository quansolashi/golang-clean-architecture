package controller

import (
	"clean-architecture/internal/application/dto"
	"clean-architecture/internal/interfaces/http/response"
	"clean-architecture/internal/interfaces/http/service"
	"clean-architecture/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     user index
// @Description list users
// @Tags        User
// @Router      /users [get]
// @Security    BearerAuth
// @Accept      json
// @Produce     json
// @Success     200 {object} response.UsersResponse
// @Failure     400 {object} util.ErrorResponse
// @Failure     500 {object} util.ErrorResponse
func (c *Controller) ListUsers(ctx *gin.Context) {
	page, err := util.GetQueryInt64(ctx, "page", 0)
	if err != nil {
		c.badRequest(ctx, err)
		return
	}
	limit := 30
	offet := (int(page) - 1) * limit

	users, err := c.Usecase.User.List(ctx, &dto.UserListParams{
		Limit:  limit,
		Offset: offet,
	})
	if err != nil {
		c.httpError(ctx, err)
		return
	}

	res := &response.UsersResponse{
		Users: service.NewUsers(users).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary     get user
// @Description detail user
// @Tags        User
// @Router      /users/{userId} [get]
// @Param				userId path uint64 true "User ID"
// @Security    BearerAuth
// @Accept      json
// @Produce     json
// @Success     200 {object} response.UserResponse
// @Failure     400 {object} util.ErrorResponse
// @Failure     404 {object} util.ErrorResponse
// @Failure     500 {object} util.ErrorResponse
func (c *Controller) GetUser(ctx *gin.Context) {
	userID, err := util.GetParamUint64(ctx, "userId", 0)
	if err != nil {
		c.badRequest(ctx, err)
		return
	}
	user, err := c.Usecase.User.Get(ctx, userID)
	if err != nil {
		c.httpError(ctx, err)
		return
	}

	res := &response.UserResponse{
		User: service.NewUser(user).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}
