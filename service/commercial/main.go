package main

import (
	"commercial-propfloor/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
	}
	// Create a Gin router
	router := gin.Default()
	router.Use(gin.Logger())
	router.Static("/css", "./css/")
	router.Static("/views", "./views/")
	router.Static("/images", "./images/")
	// Define a route to serve the index.html page

	routes.PrivateRoutes(router)

	log.Fatal(router.Run(":" + port))
}
