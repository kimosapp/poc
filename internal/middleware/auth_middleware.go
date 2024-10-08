package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kimosapp/poc/internal/core/auth"
	userRepository "github.com/kimosapp/poc/internal/core/ports/repository/users"
	"strings"
)

type AuthMiddleware struct {
	repository userRepository.Repository
}

func (midd *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}
		authorizationHeaderSplitted := strings.Split(tokenString, "Bearer ")
		if len(authorizationHeaderSplitted) != 2 {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Invalid token",
				},
			)
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(
			authorizationHeaderSplitted[1],
		)
		if err != nil {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}
		context.Set("kimosUserId", claims.ID)
		context.Next()
	}
}

func NewAuthMiddleware(repo userRepository.Repository) *AuthMiddleware {
	return &AuthMiddleware{repository: repo}
}
