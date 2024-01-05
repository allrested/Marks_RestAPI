// internal/record/service/service.go
package service

import (
	"student-api/internal/record/model"
	"student-api/internal/record/repository"
	"time"
)

// RecordService provides methods to interact with record data.
type RecordService interface {
	GetFilteredRecords(filterRequest model.FilterRequest) ([]model.FilteredRecord, error)
	SetRecords(records []model.Record) error
}

type recordService struct {
	repo repository.RecordRepository
}

// NewRecordService creates a new instance of RecordService.
func NewRecordService(repo repository.RecordRepository) RecordService {
	return &recordService{repo}
}

// GetFilteredRecords retrieves filtered records based on criteria.
func (s *recordService) GetFilteredRecords(filterRequest model.FilterRequest) ([]model.FilteredRecord, error) {
	// Convert string dates to time.Time
	startDate, err := time.Parse("2006-01-02", filterRequest.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", filterRequest.EndDate)
	if err != nil {
		return nil, err
	}

	// Call the repository method to get filtered records
	return s.repo.GetFilteredRecords(startDate, endDate, filterRequest.MinCount, filterRequest.MaxCount)
}

// SetRecords generates and inserts records into the database.
func (s *recordService) SetRecords(records []model.Record) error {
	return s.repo.SetRecords(records)
}
