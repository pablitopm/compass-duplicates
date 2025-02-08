package reader

import (
	"encoding/csv"
	"io"
	"strconv"

	"os"
	"strings"

	"main/model"

	"go.uber.org/zap"
)

type CSVReader struct {
	filePath string
	logger   *zap.Logger
}

func NewCSVReader(filePath string, logger *zap.Logger) *CSVReader {
	return &CSVReader{
		filePath: filePath,
		logger:   logger,
	}
}

func (r CSVReader) Read() ([]model.User, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []model.User
	for _, record := range records[1:] { // Skip header
		id, err := strconv.Atoi(record[0])
		if err != nil {
			r.logger.Error("Failed to convert ID to int", zap.Error(err))
			continue
		}
		user := model.User{
			ID:    id,
			Name:  record[1],
			Name1: record[2],
			Email: record[3],
		}
		user.SanitizeAddress(record[4], record[5])

		result = append(result, user)
	}
	return result, nil
}
