package routes

import (
	"dbo-be/auth"
	"dbo-be/handler"
	"dbo-be/middleware"
	"dbo-be/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, handler *handler.UserHandler, authService auth.Service, userService user.Service) {
	api.GET("/", middleware.AuthMiddleware(authService, userService), handler.GetUser)
	api.GET("/:id", middleware.AuthMiddleware(authService, userService), handler.FindUser)
	api.POST("/search", middleware.AuthMiddleware(authService, userService), handler.SearchUser)
	api.POST("/register", handler.RegisterUser)
	api.POST("/login", handler.Login)
	api.POST("/check-email", handler.CheckAvailabilityEmail)
	api.PATCH("/:id", middleware.AuthMiddleware(authService, userService), handler.EditUser)
	api.DELETE("/:id", middleware.AuthMiddleware(authService, userService), handler.DeleteUser)
}
