package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Data model
type Item struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// Database instance
var db *gorm.DB

func initDB() {
	var err error
	// Connect to SQLite
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Item{})
}

func main() {
	// Initialize database
	initDB()

	// Set up Gin
	r := gin.Default()

	// Get all items
	r.GET("/api/data", func(c *gin.Context) {
		var items []Item
		db.Find(&items)
		c.JSON(http.StatusOK, gin.H{"data": items})
	})

	// Add a new item
	r.POST("/api/data", func(c *gin.Context) {
		var newItem Item
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newItem)
		c.JSON(http.StatusOK, newItem)
	})

	// Start server
	r.Run(":8000")
}
