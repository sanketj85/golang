package controllers

import (
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	// Ensure db is initialized
	if config.DB == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	// Query users from the database and return them in JSON format
	var users []models.User

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Execute the SQL query to retrieve all users from the database and store them in the users slice.

	//rows, err := config.DB.Query("SELECT ID, FULLNAME, convert(AES_DECRYPT(PHONE,'9000')using utf8) as PHONE, convert(AES_DECRYPT(EMAIL,'9000')using utf8) as EMAIL, CITY from user where ID =?", id)
	rows, err := config.DB.Query("SELECT ID, FULLNAME, PHONE, EMAIL, CITY from user where ID =?", id)

	if err != nil {
		log.Println("Error querying users from database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying users from database"})
		return
	}
	defer rows.Close()

	// Iterate over each row and create a User struct and append it to the users slice.
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FULLNAME, &user.PHONE, &user.EMAIL, &user.CITY)
		if err != nil {
			log.Println("Error scanning user row:", err)
			continue
		}
		users = append(users, user)
	}

	// Return the users as JSON to the client.
	c.JSON(http.StatusOK, users)
}
