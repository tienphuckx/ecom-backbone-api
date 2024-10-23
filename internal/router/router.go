package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/ctl"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize the controller
	userController := ctl.NewUserController()

	// Define routes
	r.GET("/user/:id", userController.GetUser)
	r.GET("/user/email/:email", userController.GetUserByEmail)

	return r
}
