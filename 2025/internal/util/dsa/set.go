package dsa

// Set represents a set of comparable elements.
type Set[E comparable] map[E]struct{}

// NewSet creates and returns a new empty Set.
func NewSet[E comparable]() Set[E] {
	return make(Set[E])
}

// Add adds an element to the set.
func (s Set[E]) Add(element E) {
	s[element] = struct{}{}
}

// Remove deletes an element from the set.
func (s Set[E]) Remove(element E) {
	delete(s, element)
}

// Contains checks if an element exists in the set.
func (s Set[E]) Contains(element E) bool {
	_, exists := s[element]
	return exists
}

// Size returns the number of elements in the set.
func (s Set[E]) Size() int {
	return len(s)
}

// ToSlice converts the set to a slice of its elements.
func (s Set[E]) ToSlice() []E {
	slice := make([]E, 0, len(s))
	for element := range s {
		slice = append(slice, element)
	}
	return slice
}
