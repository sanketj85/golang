package controllers

import (
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	// Extract user ID from the URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user exists in the database
	validatedUser, exists := c.Get("validatedUser")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve validated user from context"})
		return
	}

	// Parse the validated user from the context
	user, ok := validatedUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse validated user from context"})
		return
	}

	// Update specific fields in the database
	if user.FULLNAME != "" {
		if !isValidName(user.FULLNAME) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name must contain only alphabetic characters"})
			return
		}
		updateFullName(id, user.FULLNAME, c)
	} else if user.PHONE != "" {
		if !isValidPhoneNumber(user.PHONE) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must be 10 digits and contain only numeric characters"})
			return
		}
		updatePhoneNumber(id, user.PHONE, c)
	} else if user.EMAIL != "" {
		if !isValidEmail(user.EMAIL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email address is not in a valid format"})
			return
		}
		updateEmail(id, user.EMAIL, c)
	} else if user.CITY != "" {
		updateCity(id, user.CITY, c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields provided to update"})
		return
	}

	// Fetch the updated user from the database
	user.ID = id
	c.JSON(http.StatusOK, user)
}

func updateFullName(id int, fullName string, c *gin.Context) {
	_, err := config.DB.Exec("UPDATE user SET FULLNAME=? WHERE ID=?", fullName, id)
	if err != nil {
		log.Println("Error updating user's full name in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user's full name in database"})
		return
	}
}

func updatePhoneNumber(id int, phoneNumber string, c *gin.Context) {
	_, err := config.DB.Exec("UPDATE user SET PHONE=? WHERE ID=?", phoneNumber, id)
	if err != nil {
		log.Println("Error updating user's phone number in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user's phone number in database"})
		return
	}
}

func updateEmail(id int, email string, c *gin.Context) {
	//_, err := config.DB.Exec("UPDATE user SET EMAIL = AES_ENCRYPT(?, '9000') WHERE ID=?", email, id)
	_, err := config.DB.Exec("UPDATE user SET EMAIL = ? WHERE ID=?", email, id)
	if err != nil {
		log.Println("Error updating user's email in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user's email in database"})
		return
	}
}

func updateCity(id int, city string, c *gin.Context) {
	_, err := config.DB.Exec("UPDATE user SET CITY=? WHERE ID=?", city, id)
	if err != nil {
		log.Println("Error updating user's city in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user's city in database"})
		return
	}
}
