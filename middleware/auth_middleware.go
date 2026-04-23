package middleware

import (
	"os"
	"strings"

	"go-framework-learing/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, "Authorization header missing", 401)
			c.Abort()
			return
		}

		// 🔥 Remove "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			utils.Error(c, "Invalid token format", 401)
			c.Abort()
			return
		}

		tokenString := parts[1]

		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// 🔥 Check signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			utils.Error(c, "Invalid token", 401)
			c.Abort()
			return
		}

		// 🔥 Safe claims extraction
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.Error(c, "Invalid token claims", 401)
			c.Abort()
			return
		}

		userID := claims["user_id"]

		// Mongo → usually string (ObjectID)
		c.Set("user_id", userID)

		c.Next()
	}
}