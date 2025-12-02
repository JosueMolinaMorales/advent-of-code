package util

func EuclideanMod(a, b int) int {
	return (a%b + b) % b
}
