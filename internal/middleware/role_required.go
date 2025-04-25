package middleware

import (
	"github.com/asylzhan/go-asylzhan-project/internal/db"
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleRequired(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in context"})
			c.Abort()
			return
		}

		var user models.User

		if err := db.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		if user.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to perform this action"})
			c.Abort()
			return
		}

		c.Next()
	}
}
