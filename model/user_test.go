package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetScore(t *testing.T) {
	tests := []struct {
		name          string
		user          User
		userToCompare User
		expectedScore int
	}{
		{
			name: "identical users",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			expectedScore: 18, // 3 (names) + 10 (email) + 5 (address)
		},
		{
			name: "different users",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "Jane",
				Name1:   "Smith",
				Email:   "jane.smith@example.com",
				Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 0,
		},
		{
			name: "some coincidence users 1",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "Jane",
				Name1:   "Smith",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 3, // 3 (email)
		},
		{
			name: "some coincidence users 2",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "Doe",
				Name1:   "John",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 6, // 3 (name) + 3 (email)
		},
		{
			name: "some coincidence users 3",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "D",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 5, // 2 (name) + 3 (email)
		},
		{
			name: "some coincidence users 4",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 4, // 1 (name) + 3 (email)
		},
		{
			name: "some coincidence users 5",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "12345", Street: "Other St", Apt: "Apt 2"},
			},
			expectedScore: 6, // 1 (name) + 3 (email) + 2 (address)
		},
		{
			name: "some coincidence users 6",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 2"},
			},
			expectedScore: 8, // 1 (name) + 3 (email) + 4 (address)
		},
		{
			name: "some coincidence users 7",
			user: User{
				ID:      1,
				Name:    "John",
				Name1:   "",
				Email:   "john.doe@example.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 2"},
			},
			userToCompare: User{
				ID:      2,
				Name:    "John",
				Name1:   "Doe",
				Email:   "john.doe@testing.com",
				Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 2"},
			},
			expectedScore: 9, // 1 (name) + 3 (email) + 5 (address)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := tt.user.GetScore(tt.userToCompare)
			assert.Equal(t, tt.expectedScore, score)
		})
	}
}

func TestUser_compareNamesScore(t *testing.T) {
	tests := []struct {
		name          string
		user          User
		userToCompare User
		expectedScore int
	}{
		{
			name:          "identical names",
			user:          User{Name: "John", Name1: "Doe"},
			userToCompare: User{Name: "John", Name1: "Doe"},
			expectedScore: 3,
		},
		{
			name:          "swapped names",
			user:          User{Name: "John", Name1: "Doe"},
			userToCompare: User{Name: "Doe", Name1: "John"},
			expectedScore: 3,
		},
		{
			name:          "partial match with initial",
			user:          User{Name: "John", Name1: "Doe"},
			userToCompare: User{Name: "John", Name1: "David"},
			expectedScore: 2, // 1 for name match + 1 for initial match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := tt.user.compareNamesScore(tt.userToCompare)
			if score != tt.expectedScore {
				t.Errorf("compareNamesScore() = %d, want %d", score, tt.expectedScore)
			}
		})
	}
}

func TestUser_compareEmailScore(t *testing.T) {
	tests := []struct {
		name          string
		user          User
		userToCompare User
		expectedScore int
	}{
		{
			name:          "identical emails",
			user:          User{Email: "john.doe@example.com"},
			userToCompare: User{Email: "john.doe@example.com"},
			expectedScore: 10,
		},
		{
			name:          "same local part, different domain",
			user:          User{Email: "john.doe@example.com"},
			userToCompare: User{Email: "john.doe@other.com"},
			expectedScore: 3,
		},
		{
			name:          "different emails",
			user:          User{Email: "john.doe@example.com"},
			userToCompare: User{Email: "jane.smith@example.com"},
			expectedScore: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := tt.user.compareEmailScore(tt.userToCompare)
			if score != tt.expectedScore {
				t.Errorf("compareEmailScore() = %d, want %d", score, tt.expectedScore)
			}
		})
	}
}

func TestUser_compareAddressScore(t *testing.T) {
	tests := []struct {
		name          string
		user          User
		userToCompare User
		expectedScore int
	}{
		{
			name:          "identical addresses",
			user:          User{Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"}},
			userToCompare: User{Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"}},
			expectedScore: 5, // 2 (street) + 1 (apt) + 2 (zip)
		},
		{
			name:          "same street, different apt",
			user:          User{Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"}},
			userToCompare: User{Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 2"}},
			expectedScore: 4, // 2 (street) + 2 (zip)
		},
		{
			name:          "different addresses",
			user:          User{Address: Address{ZipCode: "12345", Street: "Main St", Apt: "Apt 1"}},
			userToCompare: User{Address: Address{ZipCode: "67890", Street: "Other St", Apt: "Apt 2"}},
			expectedScore: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := tt.user.compareAddressScore(tt.userToCompare)
			if score != tt.expectedScore {
				t.Errorf("compareAddressScore() = %d, want %d", score, tt.expectedScore)
			}
		})
	}
}

func TestUser_SanitizeAddress(t *testing.T) {
	tests := []struct {
		name           string
		user           User
		zipcode        string
		address        string
		expectedStreet string
		expectedApt    string
	}{
		{
			name:           "simple address",
			user:           User{},
			zipcode:        "12345",
			address:        "Apt 1-Main St",
			expectedStreet: "Main St",
			expectedApt:    "Apt 1",
		},
		{
			name:           "address with P.O. Box",
			user:           User{},
			zipcode:        "67890",
			address:        "P.O. Box 123, Elm St",
			expectedStreet: " Elm St",
			expectedApt:    "P.O. Box 123",
		},
		{
			name:           "address with dots and commas",
			user:           User{},
			zipcode:        "54321",
			address:        "Apt 2-Main St., Apt 2",
			expectedStreet: "Main St Apt 2",
			expectedApt:    "Apt 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.SanitizeAddress(tt.zipcode, tt.address)
			if tt.user.Address.ZipCode != tt.zipcode {
				t.Errorf("SanitizeAddress() ZipCode = %s, want %s", tt.user.Address.ZipCode, tt.zipcode)
			}
			if tt.user.Address.Street != tt.expectedStreet {
				t.Errorf("SanitizeAddress() Street = %s, want %s", tt.user.Address.Street, tt.expectedStreet)
			}
			if tt.user.Address.Apt != tt.expectedApt {
				t.Errorf("SanitizeAddress() Apt = %s, want %s", tt.user.Address.Apt, tt.expectedApt)
			}
		})
	}
}
