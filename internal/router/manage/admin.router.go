package manage

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

// InitAdminRouter initializes admin routes
func (s *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin")
	{
		adminRouter.POST("/login", adminLoginHandler)
		adminRouter.POST("/active_user", adminActivateUserHandler)
		adminRouter.POST("/add_shop", adminAddShopHandler)
	}
}

// Handler functions for admin routes (for demonstration)
func adminLoginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Admin logged in"})
}

func adminActivateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "User activated"})
}

func adminAddShopHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Shop added"})
}
