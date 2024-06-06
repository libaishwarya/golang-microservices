package middleware

import (
	"go-gin/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckEmailExistence() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email string `json:"email" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			c.Abort()
			return
		}

		if input.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			c.Abort()
			return
		}

		var count int
		err := store.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", input.Email).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			c.Abort()
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			c.Abort()
			return
		}

		c.Next()
	}
}
