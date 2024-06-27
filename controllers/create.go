package controllers

import (
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Check if the user already exists in the database
	validatedUser, exists := c.Get("validatedUser")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve validated user from context"})
		return
	}

	// Parse the validated user from the context
	users, ok := validatedUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse validated user from context"})
		return
	}
	
	//validations
	if !isValidName(users.FULLNAME) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name must contain only alphabetic characters"})
		return
	}
	if !isValidPhoneNumber(users.PHONE) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must be 10 digits and contain only numeric characters"})
		return
	}
	if !isValidEmail(users.EMAIL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email address is not in a valid format"})
		return
	}

	// Insert into database
	result, err := config.DB.Exec("INSERT INTO user (ID,FULLNAME, PHONE, EMAIL, CITY) VALUES (?, ?, ?, ?, ?)", users.ID, users.FULLNAME, users.PHONE, users.EMAIL, users.CITY)
	
	//result, err := config.DB.Exec("INSERT INTO user (ID, FULLNAME, PHONE, EMAIL, CITY)  VALUES (?, ?, AES_ENCRYPT(?, '9000'), AES_ENCRYPT(?, '9000'), ?)", users.ID, users.FULLNAME, users.PHONE, users.EMAIL, users.CITY)

	if err != nil {
		log.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database"})
		return
	}

	// Get the inserted user's ID and set it in the response body
	id, _ := result.LastInsertId()
	users.ID = int(id)

	// Return the created user
	c.JSON(http.StatusCreated, users)
}
