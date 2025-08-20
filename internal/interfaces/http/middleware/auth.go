package middleware

import (
	"clean-architecture/internal/domain/service/auth"
	"clean-architecture/internal/infrastructure/auth/jwt"
	"clean-architecture/internal/util"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	TokenService auth.TokenService
}

func NewAuthMiddleware(tokenService auth.TokenService) *AuthMiddleware {
	return &AuthMiddleware{TokenService: tokenService}
}

func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			am.unauthorized(c, fmt.Errorf("missing authorization header"))
			return
		}

		tokenParts := strings.Split(token, "Bearer ")
		if len(tokenParts) != 2 || tokenParts[1] == "" {
			am.unauthorized(c, fmt.Errorf("invalid authorization header format"))
			return
		}

		tokenString := tokenParts[1]
		claims, err := am.TokenService.VerifyToken(c.Request.Context(), tokenString)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				am.unauthorized(c, fmt.Errorf("token has expired"))
				return
			}

			am.unauthorized(c, fmt.Errorf("authentication failed: %s", err.Error()))
			return
		}

		// Set user context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func (am *AuthMiddleware) httpError(ctx *gin.Context, err error) {
	res, code := util.NewErrorResponse(err)
	ctx.AbortWithStatusJSON(code, res)
}

func (am *AuthMiddleware) unauthorized(ctx *gin.Context, err error) {
	am.httpError(ctx, fmt.Errorf("%w: %s", util.ErrUnauthorized, err.Error()))
}
