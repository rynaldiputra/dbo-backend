package routes

import (
	"dbo-be/auth"
	"dbo-be/handler"
	"dbo-be/middleware"
	"dbo-be/user"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(api *gin.RouterGroup, handler *handler.OrderHandler, authService auth.Service, userService user.Service) {
	api.POST("/search", middleware.AuthMiddleware(authService, userService), handler.Search)
	api.GET("/", middleware.AuthMiddleware(authService, userService), handler.Get)
	api.GET("/:id", middleware.AuthMiddleware(authService, userService), handler.Find)
	api.POST("/", middleware.AuthMiddleware(authService, userService), handler.Create)
	api.PATCH("/:id", middleware.AuthMiddleware(authService, userService), handler.Edit)
	api.DELETE("/:id", middleware.AuthMiddleware(authService, userService), handler.Delete)
}
