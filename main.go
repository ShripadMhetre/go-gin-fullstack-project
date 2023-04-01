package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	helper "github.com/shripadmhetre/go-gin-fullstack-project/helpers"
	"github.com/shripadmhetre/go-gin-fullstack-project/routes"
)

func init() {
	helper.LoadEnvVariables()
}

func main() {

	router := gin.Default()

	// Setting up basic cors config
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// Routing paths
	router.GET("/dummy", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Unprotected dummy route!!!"})
	})

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Running the server
	port := helper.GetEnvVar("PORT")

	if port == "" {
		port = "8000"
	}

	router.Run(":" + port)
}
