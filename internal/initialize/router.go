package initialize

import "github.com/gin-gonic/gin"

// InitRouter initializes and returns a new Gin router
func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	v1.GET("/example", func(c *gin.Context) {
		c.String(200, "Hello from /v1/2024/example")
	})
	return r
}
