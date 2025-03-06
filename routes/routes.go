package routes

import (
	"user-management-app/controllers"
	"user-management-app/middleware"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.Login)
		api.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)

		// 🔥 FIX: Added AuthMiddleware before RequireAdmin
		adminRoutes := api.Group("/admin").Use(middleware.AuthMiddleware(), middleware.RequireAdmin)
		{
			adminRoutes.GET("/users", controllers.GetUsers)

			adminRoutes.PUT("/users/:id", controllers.UpdateUserRole)
			adminRoutes.DELETE("/users/:id", controllers.DeleteUser)
		}
	}

	return router
}
