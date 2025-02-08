package writer

import (
	"io/ioutil"
	"os"
	"testing"

	"main/model"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCSVWriter_Write(t *testing.T) {
	// Create a mock logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	testDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []struct {
		name      string
		results   []model.CompareResult
		wantErr   bool
		expectLog bool
	}{
		{
			name: "write valid results",
			results: []model.CompareResult{
				{ContactIDSource: 1, ContactIDMatch: 2, Accuracy: "HIGH"},
				{ContactIDSource: 2, ContactIDMatch: 1, Accuracy: "HIGH"},
			},
			wantErr: false,
		},
		{
			name:    "write empty result set",
			results: []model.CompareResult{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use a temporary file for testing
			tempFile, err := ioutil.TempFile(testDir, "test*.csv")
			if err != nil {
				t.Fatalf("Could not create temporary file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			writer := NewCSVWriter(tempFile.Name(), logger)
			err = writer.Write(tt.results)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
