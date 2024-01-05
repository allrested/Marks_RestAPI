// internal/record/generator/generator.go
package generator

import (
	"math/rand"
	"time"
	"student-api/internal/record/model"
)

// GenerateRecords generates random records based on the specified criteria.
func GenerateRecords(startDate, endDate time.Time, minCount, maxCount int) []model.Record {
	rand.Seed(time.Now().UnixNano())

	var records []model.Record
	numRecords := rand.Intn(maxCount-minCount+1) + minCount

	for i := 0; i < numRecords; i++ {
		createdAt := generateRandomDate(startDate, endDate)
		numMarks := rand.Intn(5) + 1 // Generate between 1 and 5 marks for each record
		marks := generateRandomMarks(numMarks)

		name := generateRandomName()

		record := model.Record{
			ID:        uint(i + 1),
			Name:      name,
			CreatedAt: createdAt,
			Marks:     marks, // Pass an array of integers
		}

		records = append(records, record)
	}

	return records
}

func generateRandomDate(startDate, endDate time.Time) time.Time {
	delta := endDate.Sub(startDate)
	randomDuration := time.Duration(rand.Int63n(int64(delta)))

	return startDate.Add(randomDuration)
}

func generateRandomMarks(numMarks int) []int {
	var marks []int
	for i := 0; i < numMarks; i++ {
		marks = append(marks, rand.Intn(11) * 10) // Generate marks between 0 and 100 for each individual mark
	}
	return marks
}

func generateRandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}
	return names[rand.Intn(len(names))]
}
