package controller

import (
	"clean-architecture/internal/application/usecase"
	"clean-architecture/internal/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Usecase *usecase.Usecase
}

func NewController(usecase *usecase.Usecase) *Controller {
	return &Controller{Usecase: usecase}
}

func (c *Controller) httpError(ctx *gin.Context, err error) {
	res, code := util.NewErrorResponse(err)
	ctx.AbortWithStatusJSON(code, res)
}

func (c *Controller) badRequest(ctx *gin.Context, err error) {
	c.httpError(ctx, fmt.Errorf("%w: %s", util.ErrInvalidArgument, err.Error()))
}

func (c *Controller) unauthorized(ctx *gin.Context, err error) {
	c.httpError(ctx, fmt.Errorf("%w: %s", util.ErrUnauthorized, err.Error()))
}

func (c *Controller) forbidden(ctx *gin.Context, err error) {
	c.httpError(ctx, fmt.Errorf("%w: %s", util.ErrForbidden, err.Error()))
}

func (c *Controller) notFound(ctx *gin.Context, err error) {
	c.httpError(ctx, fmt.Errorf("%w: %s", util.ErrNotFound, err.Error()))
}

func (c *Controller) conflict(ctx *gin.Context, err error) {
	c.httpError(ctx, fmt.Errorf("%w: %s", util.ErrConflict, err.Error()))
}
