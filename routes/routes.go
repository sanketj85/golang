// routes/routes.go

package routes

import (
	"leadAPI/controllers"
	"leadAPI/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	// Apply middleware to relevant routes

	// Grouping routes
	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/", middlewares.CheckPhoneNumberExistsMiddleware(), controllers.CreateUser)
		usersRouter.PUT("/:id", middlewares.CheckPhoneNumberExistsMiddleware(), controllers.UpdateUser)
		usersRouter.DELETE("/:id", controllers.DeleteUser)
		usersRouter.GET("/", controllers.GetUsers)
	}
	router.GET("/user/:id", controllers.GetUserByID)
	router.POST("/upload", controllers.UploadFile)
	router.GET("/files", controllers.GetUploadedFiles)
	router.GET("/files/:filename", controllers.DownloadFile)

	return router
}
