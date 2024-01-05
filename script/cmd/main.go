// cmd/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	// Retrieve database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Initialize the database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	
	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automigrate the model
	db.AutoMigrate()

	// Set up the Gin router
	router := gin.Default()

	// Define routes
	router.GET("/records", getRecordsHandler)
	
	// Run the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getRecordsHandler(c *gin.Context) {
	var requestPayload struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int    `json:"minCount"`
		MaxCount  int    `json:"maxCount"`
	}

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
		return
	}

	// Build the response payload
	responsePayload := gin.H{
		"code":    0,
		"msg":     "Success",
		"records": "Dummy response",
	}

	c.JSON(http.StatusOK, responsePayload)
}
