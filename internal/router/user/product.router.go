package user

import (
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

// InitProductRouter initializes product routes
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("product")
	{
		productRouter.POST("/add", productAddHandler)
		productRouter.GET("/list", productListHandler)
	}
}

// Handler functions for product routes (for demonstration)
func productAddHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Product added"})
}

func productListHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Product list retrieved"})
}
