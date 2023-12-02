package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}
