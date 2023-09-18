package router

import (
	"log"
	"os"

	"github.com/Ticolls/gopportunities/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Initialize Router
	router := gin.Default()

	// Config cors
	router.Use(middleware.CORSMiddleware())

	//Initialize Routes
	initializeRoutes(router)

	// Setting the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	err := router.Run(":" + port)

	if err != nil {
		log.Panicf("error: %s", err)
	}
}
