package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter initializes user routes
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("/register", userRegisterHandler)
		userRouter.POST("/send_otp", userSendOtpHandler)
		userRouter.GET("/get_info", userInfoHandler)
	}
}

// Handler functions for user routes (for demonstration)
func userRegisterHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "User registered"})
}

func userSendOtpHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OTP sent"})
}

func userInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "User info retrieved"})
}
