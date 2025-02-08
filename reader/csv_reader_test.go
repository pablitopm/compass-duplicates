package reader

import (
	"os"
	"path/filepath"
	"testing"

	"main/model"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCSVReader_Read(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	testDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []struct {
		name     string
		fileName string
		expected []model.User
		wantErr  bool
	}{
		{
			name:     "valid CSV data",
			fileName: "test_files/valid_test.csv",
			expected: []model.User{
				{
					ID:    1,
					Name:  "John",
					Name1: "Doe",
					Email: "john.doe@example.com",
					Address: model.Address{
						ZipCode: "12345",
						Street:  "7890 Main St",
						Apt:     "456",
					},
				},
				{
					ID:    2,
					Name:  "Jane",
					Name1: "Smith",
					Email: "jane.smith@example.com",
					Address: model.Address{
						ZipCode: "54321",
						Street:  "123 Elm St",
						Apt:     "Apt 1",
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "invalid ID",
			fileName: "test_files/invalid_id_test.csv",
			expected: []model.User{},
			wantErr:  false, // No error returned, but processing stops
		},
		{
			name:     "empty CSV",
			fileName: "test_files/empty_test.csv",
			expected: []model.User{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(testDir, tt.fileName)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Fatalf("Test file %s does not exist", filePath)
			}

			reader := NewCSVReader(filePath, logger)
			users, err := reader.Read()

			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Len(t, users, len(tt.expected))

				for i, user := range users {
					if user.ID != tt.expected[i].ID || user.Name != tt.expected[i].Name || user.Name1 != tt.expected[i].Name1 || user.Email != tt.expected[i].Email || user.Address != tt.expected[i].Address {
						t.Errorf("User mismatch at index %d: got %+v, expected %+v", i, user, tt.expected[i])
					}
				}
			}
		})
	}
}
