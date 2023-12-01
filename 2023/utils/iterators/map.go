package iterators

// Map maps a list based on a function
func Map[T any](list []T, f func(T) T) []T {
	var result []T
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}
