package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanzhuo-github/golang-gorm-postgres/initializers"
)

var (
	server *gin.Engine
)

// 1. Load the environment variables
// 2. Create a connection pool to the database
// 3. Create a Gin router and assigned it to the server variable.
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// Create a new route group.
	// This approach enable us to group all the routes
	// that have common middlewares or the same path prefix.
	router := server.Group("/api")

	// Define a GET route. (to the /api/healthchecker endpoint)
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// Evoke Run method to attach the route to the http.Server
	// This enable the router to start listening and serving HTTP requests.
	log.Fatal(server.Run(":" + config.ServerPort))
}
