package middleware

import (
	"net/http"

	"books/token"

	"github.com/gin-gonic/gin"
)

func RequireJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
