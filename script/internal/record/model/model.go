// internal/record/model/model.go
package model

import "time"

// Record represents the basic structure of a record.
type Record struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Marks     []int     `json:"marks" gorm:"type:json"`

	CreatedAt time.Time `json:"createdAt"`
}

// FilteredRecord represents the structure for filtered records.
type FilteredRecord struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	TotalMarks  int       `json:"totalMarks"`
}

// FilterRequest represents the structure for filtering records.
type FilterRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}
