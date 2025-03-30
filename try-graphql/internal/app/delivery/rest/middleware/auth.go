package middleware

import (
	"strings"
	"try-graphql/config"
	"try-graphql/internal/app/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().JWTSecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrTokenExpired {
				c.AbortWithStatusJSON(401, entity.ErrTokenExpired)
				return
			}
			c.AbortWithStatusJSON(401, entity.ErrTokenInvalid)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["sub"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, entity.ErrTokenInvalid)
		}
	}
}
