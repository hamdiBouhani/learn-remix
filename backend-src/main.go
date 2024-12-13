package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go backend started!",
		})
	})

	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []int{1, 2, 3, 4, 5},
		})
	})

	// Start the server
	r.Run(":8000") // Runs on http://localhost:8000
}
