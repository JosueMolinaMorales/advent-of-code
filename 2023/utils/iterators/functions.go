package iterators

// Filter filters a list based on a predicate
func Filter[T any](list []T, f func(T) bool) []T {
	var filtered []T
	for _, v := range list {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Map maps a list based on a function
func Map[T any](list []T, f func(T) T) []T {
	var result []T
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}

func Product[T ~int | ~float32 | ~float64](list []T) T {
	var result T = 1
	for _, v := range list {
		result *= v
	}
	return result
}

func Every[T any](list []T, f func(T) bool) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func Some[T any](list []T, f func(T) bool) bool {
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
