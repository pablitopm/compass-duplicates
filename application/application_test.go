package application

import (
	"errors"
	"main/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestApplicationRun(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tests := []struct {
		name        string
		reader      Reader
		writer      Writer
		userService UserService
		wantErr     bool
	}{
		{
			name: "successful run",
			reader: &ReaderMock{
				ReadFunc: func() ([]model.User, error) {
					return []model.User{
						{ID: 1, Name: "John Doe", Email: ""},
						{ID: 2, Name: "Jane Doe", Email: ""},
					}, nil
				},
			},
			writer: &WriterMock{
				WriteFunc: func(result []model.CompareResult) error {
					return nil
				},
			},
			userService: &UserServiceMock{
				CompareAndClassifyFunc: func(users []model.User) []model.CompareResult {
					return []model.CompareResult{
						{ContactIDSource: 1, ContactIDMatch: 2, Accuracy: "100%"},
					}
				},
			},
		},
		{
			name: "error read",
			reader: &ReaderMock{
				ReadFunc: func() ([]model.User, error) { return nil, errors.New("some error") },
			},
			wantErr:     true,
			writer:      &WriterMock{},
			userService: &UserServiceMock{},
		},
		{
			name: "write read",
			reader: &ReaderMock{
				ReadFunc: func() ([]model.User, error) {
					return []model.User{
						{ID: 1, Name: "John Doe", Email: ""},
						{ID: 2, Name: "Jane Doe", Email: ""},
					}, nil
				},
			},
			writer: &WriterMock{
				WriteFunc: func(result []model.CompareResult) error { return errors.New("some error") },
			},
			wantErr: true,
			userService: &UserServiceMock{
				CompareAndClassifyFunc: func(users []model.User) []model.CompareResult {
					return []model.CompareResult{
						{ContactIDSource: 1, ContactIDMatch: 2, Accuracy: "100%"},
					}
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApplication(logger, tt.reader, tt.writer, tt.userService)
			err := app.Run()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
