package controllers

import (
	"fmt"
	"leadAPI/config"
	"leadAPI/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	// Validations
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
	result, err := config.DB.Exec("INSERT INTO encuser (ID, FULLNAME, PHONE, EMAIL, CITY) VALUES (?, ?, AES_ENCRYPT(?, '9000'), AES_ENCRYPT(?, '9000'), ?)", users.ID, users.FULLNAME, users.PHONE, users.EMAIL, users.CITY)
	if err != nil {
		log.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database"})
		return
	}

	// Get the inserted user's ID and set it in the response body
	id, _ := result.LastInsertId()
	users.ID = int(id)

	// Send welcome email
	if err := sendWelcomeEmail(users); err != nil {
		log.Println("Error sending welcome email:", err)
		// Handle error (return an error response or continue)
	}

	// Return the created user
	c.JSON(http.StatusCreated, users)
}

func sendWelcomeEmail(user *models.User) error {
	// Create a new SendGrid client

	// // Compose the email details
	from := mail.NewEmail("Sanket", "sanket.jadhav@choicetechlab.com")
	subject := "Welcome to Your App!"
	to := mail.NewEmail(user.FULLNAME, user.EMAIL) // User's name and email
	plainTextContent := "Dear " + user.FULLNAME + ",\n\nWelcome to Choice! We are thrilled to have you join us and embark on this journey with us.\n\nAt Choice, we are committed to serve you. We're here to support you every step of the way.\n\nOnce again, welcome aboard! We're excited to have you onboard.\n\nBest regards,\nSanket Jadhav\nChoice Techlab Solutions"
	htmlContent := "<p>Dear " + user.FULLNAME + ",</p><p>Welcome to Choice! We are thrilled to have you join us and embark on this journey with us.</p><p>At Choice, we are committed to serve you. We're here to support you every step of the way.</p><p>Once again, welcome aboard! We're excited to have you onboard.</p><p>Best regards,<br>Sanket Jadhav<br>Choice Techlab Solutions</p>"


	// Create the email message
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Send the email using the SendGrid client
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}