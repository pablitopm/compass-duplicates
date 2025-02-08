package service

import (
	"fmt"
	"testing"
	"time"

	"main/model"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CompareAndClassify(t *testing.T) {
	// Mock cache setup

	tests := []struct {
		name        string
		users       []model.User
		expected    []model.CompareResult
		cache       Cache
		expectCache []string // Keys expected to be in cache after test
	}{
		{
			name: "simple comparison cache used",
			users: []model.User{
				{ID: 1, Name: "John", Name1: "Doe", Email: "john@example.com"},
				{ID: 2, Name: "John", Name1: "Doe", Email: "john@example.com"},
			},
			expected: []model.CompareResult{
				{ContactIDSource: 1, ContactIDMatch: 2, Accuracy: "HIGH"},
				{ContactIDSource: 2, ContactIDMatch: 1, Accuracy: "HIGH"},
			},
			cache: &CacheMock{
				SetFunc: func(key string, value interface{}, duration time.Duration) {
					assert.Equal(t, "1-2", key)
				},
				GetFunc: func(key string) (interface{}, bool) {
					assert.Equal(t, "1-2", key)
					return nil, false
				},
			},
			expectCache: []string{"1-2"},
		},
		{
			name: "without cache validation",
			users: []model.User{
				{ID: 1, Name: "John", Name1: "Doe", Email: "john@example.com"},
				{ID: 2, Name: "John", Name1: "Doe", Email: "john@example.com"},
				{ID: 3, Name: "Jane", Name1: "Smith", Email: "jane@example.com"},
			},
			expected: []model.CompareResult{
				{ContactIDSource: 1, ContactIDMatch: 2, Accuracy: "HIGH"},
				{ContactIDSource: 1, ContactIDMatch: 3, Accuracy: "LOW"},
				{ContactIDSource: 2, ContactIDMatch: 1, Accuracy: "HIGH"},
				{ContactIDSource: 2, ContactIDMatch: 3, Accuracy: "LOW"},
				{ContactIDSource: 3, ContactIDMatch: 1, Accuracy: "LOW"},
				{ContactIDSource: 3, ContactIDMatch: 2, Accuracy: "LOW"},
			},
			cache: &CacheMock{
				SetFunc: func(key string, value interface{}, duration time.Duration) {},
				GetFunc: func(key string) (interface{}, bool) {
					return nil, false
				},
			},
			expectCache: []string{"1-2", "1-3", "2-3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewUserService(tt.cache)
			results := service.CompareAndClassify(tt.users)

			// Check results
			assert.Len(t, results, len(tt.expected))

			for i, result := range results {
				if result.ContactIDSource != tt.expected[i].ContactIDSource || result.ContactIDMatch != tt.expected[i].ContactIDMatch || result.Accuracy != tt.expected[i].Accuracy {
					t.Errorf("Result mismatch at index %d: got %+v, expected %+v", i, result, tt.expected[i])
				}
			}
		})
	}
}

func TestUserService_classifyScore(t *testing.T) {
	service := UserService{}
	tests := []struct {
		score    int
		expected string
	}{
		{score: 9, expected: "HIGH"},
		{score: 8, expected: "MID"},
		{score: 5, expected: "MID"},
		{score: 4, expected: "LOW"},
		{score: 0, expected: "LOW"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("score %d", tt.score), func(t *testing.T) {
			result := service.classifyScore(tt.score)
			assert.Equal(t, tt.expected, result)
		})
	}
}
