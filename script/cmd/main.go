// cmd/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"student-api/internal/record/model"
	"student-api/internal/record/handler"
	"student-api/internal/record/repository"
	"student-api/internal/record/service"
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
	db.AutoMigrate(&model.Record{})

	// Create instances of repository, service, and handler
	recordRepo := repository.NewRecordRepository(db)
	recordService := service.NewRecordService(recordRepo)
	recordHandler := handler.NewRecordHandler(recordService)

	// Set up the Gin router
	router := gin.Default()

	// Define your endpoints
	router.GET("/migrations", recordHandler.SetRecords)
	router.POST("/records", recordHandler.GetFilteredRecords)
	
	// Run the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
