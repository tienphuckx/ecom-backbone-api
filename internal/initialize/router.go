package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/router/manage"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/router/user"
)

// InitRouter initializes the Gin engine and sets up the routes
func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.SysConfig.Server.Mode == "dev" { // Run in development mode
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode) // Run in production mode
		r = gin.New()
	}

	// Set up the main group for routes, e.g., "/v1/2024"
	MainGroup := r.Group("/v1/2024")

	// Status route for monitoring
	MainGroup.GET("/checkStatus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	// Initialize user routes (for general users)
	userRouter := user.RouterGroupApp.UserRouter // Access UserRouter, not User
	userRouter.InitUserRouter(MainGroup)         // Load user routes for general users

	// Initialize product routes (optional if needed)
	productRouter := user.RouterGroupApp.ProductRouter
	productRouter.InitProductRouter(MainGroup) // Load product routes for general users

	// Initialize admin routes
	adminRouter := manage.RouterGroupApp.AdminRouter
	adminRouter.InitAdminRouter(MainGroup) // Load admin routes

	// Initialize admin's user routes (for managing users)
	adminUserRouter := manage.RouterGroupApp.UserRouter
	adminUserRouter.InitUserRouter(MainGroup) // Load user management routes for admins

	return r
}
