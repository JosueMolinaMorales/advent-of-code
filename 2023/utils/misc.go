package utils

import (
	"math/big"
	"strconv"
)

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

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Calculate the greatest common divisor (GCD) using Euclid's algorithm
func Gcd(a, b *big.Int) *big.Int {
	gcd := new(big.Int)
	return gcd.GCD(nil, nil, a, b)
}

// Calculate the least common multiple (LCM) of two numbers
func Lcm(a, b *big.Int) *big.Int {
	// LCM(a, b) = |a * b| / GCD(a, b)
	absA := new(big.Int).Abs(a)
	absB := new(big.Int).Abs(b)
	gcdAB := Gcd(absA, absB)

	// LCM = |a * b| / GCD(a, b)
	lcm := new(big.Int).Div(new(big.Int).Mul(absA, absB), gcdAB)
	return lcm
}

// Calculate the least common multiple (LCM) of multiple numbers
func MultipleLCM(numbers []int) int {
	// Initialize the LCM with 1
	result := big.NewInt(1)

	// Iterate through each number and update the LCM
	for _, num := range numbers {
		result = Lcm(result, big.NewInt(int64(num)))
	}

	return int(result.Int64())
}
