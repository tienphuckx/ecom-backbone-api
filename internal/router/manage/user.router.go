package manage

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter initializes routes for admins to manage users
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("admin/user")
	{
		userRouter.PUT("/update/:id", updateUserHandler)          // Admin updates user info
		userRouter.POST("/deactivate/:id", deactivateUserHandler) // Admin deactivates a user
		userRouter.DELETE("/delete/:id", deleteUserHandler)       // Admin deletes a user
	}
}

// Handler functions for admin-user routes

// updateUserHandler allows an admin to update user info
func updateUserHandler(c *gin.Context) {
	userId := c.Param("id")
	// Update user logic goes here (parse body and perform update)
	c.JSON(200, gin.H{"message": "User updated successfully", "userId": userId})
}

// deactivateUserHandler allows an admin to deactivate a user
func deactivateUserHandler(c *gin.Context) {
	userId := c.Param("id")
	// Deactivate user logic goes here
	c.JSON(200, gin.H{"message": "User deactivated", "userId": userId})
}

// deleteUserHandler allows an admin to delete a user
func deleteUserHandler(c *gin.Context) {
	userId := c.Param("id")
	// Delete user logic goes here
	c.JSON(200, gin.H{"message": "User deleted", "userId": userId})
}
