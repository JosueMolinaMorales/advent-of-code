package util

func EuclideanMod(a, b int) int {
	return (a%b + b) % b
}

// GCF finds the greatest common factor of two integers using recursion.
func GCF(a, b int) int {
	if b == 0 {
		return a
	}
	return GCF(b, a%b)
}
