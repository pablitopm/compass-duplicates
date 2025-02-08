package writer

import (
	"encoding/csv"
	"fmt"
	"main/model"
	"os"

	"go.uber.org/zap"
)

type CSVWriter struct {
	filePath string
	logger   *zap.Logger
}

func NewCSVWriter(filePath string, logger *zap.Logger) *CSVWriter {
	return &CSVWriter{
		filePath: filePath,
		logger:   logger,
	}
}

func (s CSVWriter) Write(result []model.CompareResult) error {
	// Create the file
	fmt.Print(s.filePath)
	file, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Create a new csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	err = writer.Write([]string{"ContactIDSource", "ContactIDMatch", "Accuracy"})
	if err != nil {
		return fmt.Errorf("error writing header to csv: %v", err)
	}

	// Write all the records
	for _, res := range result {
		record := []string{fmt.Sprintf("%d", res.ContactIDSource), fmt.Sprintf("%d", res.ContactIDMatch), res.Accuracy}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to csv: %v", err)
		}
	}
	return nil
}
