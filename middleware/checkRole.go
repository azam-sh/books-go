package middleware

import (
	"books/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckRole(allowedID uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, err := token.ExtractRoleID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if roleID > allowedID {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you do not have permission",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
