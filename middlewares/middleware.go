package middlewares

import (
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPhoneNumberExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user data from JSON request body into a User struct
		var users models.User
		if err := c.ShouldBindJSON(&users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}

		// Check if the phone number already exists
		var count int
		//err := config.DB.QueryRow("SELECT COUNT(*) FROM user WHERE PHONE = ?", users.PHONE).Scan(&count)

		err := config.DB.QueryRow("SELECT COUNT(*) FROM encuser WHERE convert(AES_DECRYPT(PHONE,'9000')using utf8) = ?", users.PHONE).Scan(&count)
		if err != nil {
			log.Println("Error checking phone number existence:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking phone number existence"})
			c.Abort()
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already exists"})
			c.Abort()
			return
		} else {

			// If phone number does not exist, store validated user in context
			c.Set("validatedUser", &users)
			c.Next()
		}
	}
}
