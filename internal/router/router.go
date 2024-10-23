package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/ctl"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/middlewares"
)

// Middleware AA
func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next() // Calls the next handler/middleware in the chain
		fmt.Println("After --> AA")
	}
}

// Middleware BB
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next() // Calls the next handler/middleware in the chain
		fmt.Println("After --> BB")
	}
}

// Middleware CC
func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> CC")
		c.Next() // Calls the next handler/middleware in the chain
		fmt.Println("After --> CC")
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize the controller
	userController := ctl.NewUserController()

	// Attach middleware globally (for every route)
	r.Use(middlewares.AuthenticateMiddleware(), AA(), BB(), CC())

	// Define routes
	r.GET("/user/:id", userController.GetUser)
	r.GET("/user/email/:email", userController.GetUserByEmail)

	return r
}

func main() {
	// Initialize the router
	r := NewRouter()

	// Define a route group with middleware attached to it
	v1 := r.Group("/v1/2024")
	v1.GET("/example", func(c *gin.Context) {
		c.String(200, "Hello from /v1/2024/example")
	})

	// Start the server
	r.Run(":8080")
}

// NewRouter sets up the router and middleware
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Initialize the controller
	userController := ctl.NewUserController()

	// Attach middleware globally (for every route)
	r.Use(AA(), BB(), CC())

	// Define routes
	r.GET("/user/:id", userController.GetUser)
	r.GET("/user/email/:email", userController.GetUserByEmail)
	return r
}
