package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shripadmhetre/go-gin-fullstack-project/controllers"
	middleware "github.com/shripadmhetre/go-gin-fullstack-project/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
