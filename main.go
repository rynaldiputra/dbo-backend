package main

import (
	"dbo-be/auth"
	"dbo-be/config"
	"dbo-be/handler"
	"dbo-be/order"
	"dbo-be/routes"
	"dbo-be/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfig()
	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database sukses terkoneksi")

	err = config.Migrate()

	if err != nil {
		log.Fatal(err.Error())
	}

	// call repository
	userRepository := user.NewRepository(db)
	orderRepository := order.NewRepository(db)

	// call service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	orderService := order.NewService(orderRepository)

	// call handler
	userHandler := handler.NewUserHandler(userService, authService)
	orderHandler := handler.NewOrderHandler(orderService)

	// gin router
	router := gin.Default()

	// api versioning
	userApi := router.Group("/api/v1/user")
	orderApi := router.Group("/api/v1/order")

	routes.UserRoutes(userApi, userHandler, authService, userService)
	routes.OrderRoutes(orderApi, orderHandler, authService, userService)

	router.Run()
}
