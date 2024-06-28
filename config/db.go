// config/init.go

package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver for Go
)

var DB *sql.DB

// initialize the database connection
func InitDB(dataSourceName string) {
	var err error
	// Open a database connection using data source name
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection to make sure it's working
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
}
