package main

import (
	"fmt"
	"log"

	"github.com/hanzhuo-github/golang-gorm-postgres/initializers"
	"github.com/hanzhuo-github/golang-gorm-postgres/models"
)

// 1. load the environment variables
// 2. create a connection pool to the database
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	fmt.Print(config)

	initializers.ConnectDB(&config)
}

// call the AutoMigrate() function provided by GORM
// to create the database migrate
// and push changes to the database.
func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
