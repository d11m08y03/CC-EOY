package routes

import (
	"github.com/d11m08y03/CC-EOY/controllers"
	"github.com/d11m08y03/CC-EOY/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/login", controllers.Login)

	// Route only available during testing
	r.POST("/create-admin", controllers.CreateAdmin)

	// Admin-only route for registration
	r.POST("/register", middleware.AdminAuthMiddleware(), controllers.Register)

	// Protected routes
	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.JWTAuthMiddleware())
	{
		authRoutes.POST("/students", controllers.CreateStudent)
		authRoutes.PUT("/students", controllers.MarkStudentAsPresent)
	}

	return r
}
