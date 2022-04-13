package middlewares_auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/garcia-paulo/go-gin/application/token"
	"github.com/garcia-paulo/go-gin/application/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenMaker *token.TokenMaker
}

func NewAuthMiddleware(tokenMaker *token.TokenMaker) *AuthMiddleware {
	return &AuthMiddleware{
		tokenMaker: tokenMaker,
	}
}

func (a *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("authorization")
		if len(header) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(fmt.Errorf("authorization header not provided")))
			return
		}

		fields := strings.Fields(header)
		if len(fields) < 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(fmt.Errorf("invalid authorization format")))
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != "bearer" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(fmt.Errorf("unsupported authorization type %s", authType)))
			return
		}

		token := fields[1]
		payload, err := a.tokenMaker.VerifyToken(token)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		context.Set("authorization_payload", payload)
		context.Next()
	}
}
