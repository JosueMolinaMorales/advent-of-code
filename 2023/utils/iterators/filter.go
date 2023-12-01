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
