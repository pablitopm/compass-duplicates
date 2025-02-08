package model

import (
	"strings"
	"sync"
)

type User struct {
	ID      int
	Name    string
	Name1   string
	Email   string
	Address Address
}

type Address struct {
	ZipCode string
	Street  string
	Apt     string
}

func (u User) GetScore(userToCompare User) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	score := 0

	// Start three goroutines to calculate scores concurrently
	wg.Add(3)
	go func() {
		defer wg.Done()
		mu.Lock()
		score += u.compareNamesScore(userToCompare)
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		score += u.compareEmailScore(userToCompare)
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		score += u.compareAddressScore(userToCompare)
		mu.Unlock()
	}()

	wg.Wait()

	return score
}

// This functions compares name and last name of two users
func (u User) compareNamesScore(userToCompare User) int {
	// names match, no need to look further
	if u.Name == userToCompare.Name && u.Name1 == userToCompare.Name1 {
		return 3
	}
	// this check is because input provided declared name and name1 as columns, not name and lastname
	// names match from different fields, no need to look further
	if u.Name == userToCompare.Name1 && u.Name1 == userToCompare.Name {
		return 3
	}

	score := 0
	if u.Name == userToCompare.Name {
		score += 1
		//check if initial of name1 is the same
		if u.Name1 != "" && userToCompare.Name1 != "" && u.Name1[0] == userToCompare.Name1[0] {
			score += 1
		}
	}

	if u.Name1 == userToCompare.Name1 {
		score += 1
		//check if initial of name is the same
		if u.Name != "" && userToCompare.Name != "" && u.Name[0] == userToCompare.Name[0] {
			score += 1
		}

	}

	return score
}

// This functions compares Email
func (u User) compareEmailScore(userToCompare User) int {
	// no email to compare
	if u.Email == "" || userToCompare.Email == "" {
		return 0
	}

	sanitizeEmail := func(email string) string {
		return strings.ToLower(strings.TrimSpace(email))
	}
	email1 := sanitizeEmail(u.Email)
	email2 := sanitizeEmail(userToCompare.Email)

	//same email, same person
	if email1 == email2 {
		return 10
	}

	//same email but distinct @
	if strings.Split(email1, "@")[0] == strings.Split(email2, "@")[0] {
		return 3
	}

	return 0
}

// This functions compares Address (zip and address)
func (u User) compareAddressScore(userToCompare User) int {
	score := 0
	// we compare streets and apt/house number
	if u.Address.Street != "" && userToCompare.Address.Street != "" && u.Address.Street == userToCompare.Address.Street {
		score += 2
		if u.Address.Apt == userToCompare.Address.Apt {
			score += 1
		}
	}

	if u.Address.ZipCode == userToCompare.Address.ZipCode {
		score += 2
	}

	return score
}

// This creates a struct for Address to have better encapsulation and understanding of what address is
func (u *User) SanitizeAddress(zipcode, address string) {
	street, apt := parseAddress(address)
	// removing extra dots in street to sanitize better
	street = strings.ReplaceAll(street, ".", "")
	street = strings.ReplaceAll(street, ",", "")
	u.Address = Address{
		ZipCode: zipcode,
		Street:  street,
		Apt:     apt,
	}
}

// parseAddress takes a string address and splits it into Street, Apt, and P.O. Box if present.
func parseAddress(addr string) (street, apt string) {
	// Trim any surrounding quotes from the address string.
	addr = strings.Trim(addr, "\"")

	// Check if there's an apartment or P.O. Box number.
	if strings.Contains(addr, "Ap #") {
		parts := strings.Split(addr, "-")
		return parts[1], parts[0]
	} else if strings.Contains(addr, "P.O. Box") {
		parts := strings.Split(addr, ",")
		return parts[1], parts[0]
	} else {
		parts := strings.Split(addr, "-")
		if len(parts) > 1 {
			return parts[1], parts[0]
		}
	}
	return addr, ""
}
