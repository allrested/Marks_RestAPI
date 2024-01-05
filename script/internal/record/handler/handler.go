// internal/record/handler/handler.go
package handler

import (
	"github.com/gin-gonic/gin"
	"student-api/internal/record/generator"
	"student-api/internal/record/model"
	"student-api/internal/record/service"
	"time"
)

type RecordHandler struct {
	recordService service.RecordService
}

// NewRecordHandler creates a new instance of RecordHandler.
func NewRecordHandler(recordService service.RecordService) *RecordHandler {
	return &RecordHandler{recordService}
}

// SetRecords generates and inserts records into the database.
func (h *RecordHandler) SetRecords(c *gin.Context) {
	// Generate random records
	startDate, _ := time.Parse("2006-01-02", "2016-01-26")
	endDate, _ := time.Parse("2006-01-02", "2018-02-02")
	records := generator.GenerateRecords(startDate, endDate, 100, 300)

	// Insert generated records into the database
	err := h.recordService.SetRecords(records)
	if err != nil {
		c.JSON(500, gin.H{"code": 2, "msg": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "Success"})
}

func (h *RecordHandler) GetFilteredRecords(c *gin.Context) {
	// Parse JSON payload from the request
	var filterRequest model.FilterRequest
	if err := c.ShouldBindJSON(&filterRequest); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Bad Request"})
		return
	}

	// Call the service method to get filtered records
	filteredRecords, err := h.recordService.GetFilteredRecords(filterRequest)
	if err != nil {
		c.JSON(500, gin.H{"code": 2, "msg": "Internal Server Error"})
		return
	}

	// Respond with the data
	c.JSON(200, gin.H{"code": 0, "msg": "Success", "records": filteredRecords})
}
