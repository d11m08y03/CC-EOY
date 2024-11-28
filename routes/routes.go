package routes

import (
	"github.com/d11m08y03/CC-EOY/controllers"
	"github.com/d11m08y03/CC-EOY/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

  // Protected routes
	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.JWTAuthMiddleware())
	{
		authRoutes.GET("/profile", func(c *gin.Context) {
			userID := c.MustGet("user_id")
			email := c.MustGet("email")
			c.JSON(200, gin.H{"user_id": userID, "email": email})
		})
	}

	return r
}
