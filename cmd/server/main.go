package main

import (
	"github.com/tienphuckx/ecom-backbone-api.git/internal/router"
)

func main() {
	// Set up routes
	r := router.SetupRouter()

	// Start the server on port 8080
	r.Run(":8080")
}
