// internal/record/repository/repository.go
package repository

import (
	"encoding/json"
	"gorm.io/gorm"
	"student-api/internal/record/model"
	"time"
)

// RecordRepository handles the database operations for records.
type RecordRepository interface {
	GetFilteredRecords(startDate, endDate time.Time, minCount, maxCount int) ([]model.FilteredRecord, error)
	SetRecords(records []model.Record) error
}

type recordRepository struct {
	db *gorm.DB
}

// NewRecordRepository creates a new instance of RecordRepository.
func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db}
}

// GetFilteredRecords retrieves filtered records based on criteria.
func (r *recordRepository) GetFilteredRecords(startDate, endDate time.Time, minCount, maxCount int) ([]model.FilteredRecord, error) {
	var filteredRecords []model.FilteredRecord

	err := r.db.Raw(`
	SELECT
		id,
		created_at,
		JSON_UNQUOTE(SUM(value)) AS total_marks
	FROM
		records,
		JSON_TABLE(marks, "$[*]" COLUMNS (value INT PATH "$")) AS jt
	WHERE
		created_at BETWEEN ? AND ?
	GROUP BY
		id, created_at
	HAVING
		total_marks BETWEEN ? AND ?
	`, startDate, endDate, minCount, maxCount).Scan(&filteredRecords).Error

	if err != nil {
		return nil, err
	}

	return filteredRecords, nil
}

// SetRecords generates and inserts records into the database.
func (r *recordRepository) SetRecords(records []model.Record) error {
	
	// Drop the table if it exists
	if err := r.db.Migrator().DropTable(&model.Record{}); err != nil {
		return err
	}

	// AutoMigrate the model
	if err := r.db.AutoMigrate(&model.Record{}); err != nil {
		return err
	}
	
	for _, record := range records {
		// Convert the Marks array to a JSON string
		marksJSON, err := json.Marshal(record.Marks)
		if err != nil {
			return err
		}

		// Use raw SQL to insert the record into the database with the JSON array
		if err := r.db.Exec("INSERT INTO records (name, marks, created_at, id) VALUES (?, ?, ?, ?)",
			record.Name, string(marksJSON), record.CreatedAt, record.ID).Error; err != nil {
			return err
		}
	}
	return nil
}