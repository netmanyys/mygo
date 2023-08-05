package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Serve static files from the "static" directory (including the index.html file)
	r.StaticFile("/", "./index.html")

	// Run the server on port 9999
	r.Run(":9999")
}
