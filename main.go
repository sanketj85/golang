// main.go

package main

import (
	"leadAPI/config"
	"leadAPI/routes"
)

func main() {
	// Initialize database connection
	config.InitDB("root:pass@123@tcp(localhost:3306)/trainingdb")

	// Setup router
	router := routes.SetupRouter()

	// Run the server
	router.Run(":8000")
}
