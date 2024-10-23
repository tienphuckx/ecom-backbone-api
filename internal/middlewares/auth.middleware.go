package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/response"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization token from the request header
		token := c.GetHeader("Authorization")

		// Validate the token (this is a mock check; replace it with real token validation)
		if token != "valid-token" {
			// Respond with an invalid token error and stop further request processing
			response.ServerResponseError(c, response.ErrCode_INVALID_TOKEN, "Invalid token", nil)
			c.Abort() // Stop further middleware/handler processing
			return
		}

		// Token is valid; proceed to the next handler/middleware
		c.Next()
	}
}

// AuthorizeMiddleware checks if the user has the required role to access the resource
func AuthorizeMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Extract user's role from the request header or context
		// Get the user's role from the request context (in a real app, you'd get it from the token or session)
		userRole := c.GetHeader("Role") // Mock role retrieval; replace with actual logic

		// Check if the user has the required role
		if userRole != requiredRole {
			// Respond with an unauthorized access error and stop further request processing
			response.ServerResponseError(c, response.ErrCode_FORBIDDEN,
				"You do not have permission to access this resource", nil)
			c.Abort() // Stop further middleware/handler processing
			return
		}

		// User has the required role; proceed to the next handler/middleware
		c.Next()
	}
}
