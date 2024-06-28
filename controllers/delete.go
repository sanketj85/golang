package controllers

import (
	"leadAPI/config"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Extract user ID from the URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Delete from database
	_, err = config.DB.Exec("DELETE FROM encuser WHERE ID=?", id)
	if err != nil {
		log.Println("Error deleting user from database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user from database"})
		return
	}

	// Return a success message to the client
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
