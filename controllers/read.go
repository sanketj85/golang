package controllers

import (
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Ensure db is initialized
	if config.DB == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	// Extracting Pagination parameters from the URL
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	// Parse page and pageSize parameters and handle invalid inputs
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // default to page 1 if invalid or not provided
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 5 // default page size if invalid or not provided
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Query users from the database with pagination
	var users []models.User

	// Execute the SQL query with pagination and store the result in the users slice.
	query := "SELECT ID, FULLNAME, PHONE, EMAIL, CITY FROM user LIMIT ? OFFSET ?"
	//query := "SELECT ID, FULLNAME, convert(AES_DECRYPT(PHONE,'9000')using utf8) as PHONE, convert(AES_DECRYPT(EMAIL,'9000')using utf8) as EMAIL, CITY from user LIMIT? OFFSET?"
	rows, err := config.DB.Query(query, pageSize, offset)
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

	// Check for errors during rows iteration
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over user rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over user rows"})
		return
	}

	// Return the users and pagination metadata as JSON to the client.
	c.JSON(http.StatusOK, gin.H{
		"users":      users,
		"page":       page,
		"pageSize":   pageSize,
		"totalCount": len(users), 
	})
}
