package utils

import "strconv"

// IsDigit checks if a string is a digit
func IsDigit(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

// FindAllIndex returns a slice of indices where a substring is found
func FindAllIndex(s string, substr string) []int {
	var indices []int
	for i := 0; i < len(s); i++ {
		if len(s[i:]) < len(substr) {
			break
		}
		if s[i:i+len(substr)] == substr {
			indices = append(indices, i)
		}
	}
	return indices
}
