package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/auth") {
			ctx.Next()
			return
		}
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised authentication header"})
			ctx.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
