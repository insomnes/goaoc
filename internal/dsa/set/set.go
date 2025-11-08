package set

import "fmt"

// Set represents a collection of unique elements of type T.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates and returns a new Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// NewSetCap preallocates buckets to reduce rehashing on many inserts.
func NewSetCap[T comparable](capHint int) *Set[T] {
	if capHint < 0 {
		capHint = 0
	}
	return &Set[T]{elements: make(map[T]struct{}, capHint)}
}

// FromSlice creates a new Set from a slice of elements.
func FromSlice[T comparable](elements []T) *Set[T] {
	set := NewSet[T]()
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// FromString creates a new Set from a string, treating each character as an element.
func FromString(s string) *Set[rune] {
	set := NewSet[rune]()
	for _, char := range s {
		set.Add(char)
	}
	return set
}

// FromIter creates a new Set from an iterator function.
func FromIter[T comparable](iter func(yield func(T) bool)) *Set[T] {
	set := NewSet[T]()
	iter(func(element T) bool {
		set.Add(element)
		return true
	})
	return set
}

// Add inserts an element into the set.
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Remove deletes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if an element is in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// IsEmpty checks if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Elements returns a slice containing all elements in the set.
func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

// Iter returns an iterator over the set's elements.
func (s *Set[T]) Iter() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for key := range s.elements {
			if !yield(key) {
				return
			}
		}
	}
}

// Copy creates a shallow copy of the set.
func (s *Set[T]) Copy() *Set[T] {
	newSet := NewSet[T]()
	for key := range s.elements {
		newSet.Add(key)
	}
	return newSet
}

// Union returns a new set that is the union of the current set and another set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := s.Copy()
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

// Intersection returns a new set that is the intersection of the current set and another set.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	toIter, toCheck := s, other
	if other.Size() < s.Size() {
		toIter, toCheck = other, s
	}
	toIter.Iter()(func(element T) bool {
		if toCheck.Contains(element) {
			result.Add(element)
		}
		return true
	})
	return result
}

// Difference returns a new set that is the difference of the current set and another set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	s.Iter()(func(element T) bool {
		if !other.Contains(element) {
			result.Add(element)
		}
		return true
	})
	return result
}

// IsSubset checks if the current set is a subset of another set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.elements {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the current set is a superset of another set.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equal checks if the current set is equal to another set
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.elements {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// ClearIf removes elements from the set that satisfy the given predicate.
func (s *Set[T]) ClearIf(predicate func(T) bool) {
	for key := range s.elements {
		if predicate(key) {
			delete(s.elements, key)
		}
	}
}

// RetainIf retains only the elements in the set that satisfy the given predicate.
func (s *Set[T]) RetainIf(predicate func(T) bool) {
	for key := range s.elements {
		if !predicate(key) {
			delete(s.elements, key)
		}
	}
}

// AddAll inserts all provided elements.
func (s *Set[T]) AddAll(elements ...T) {
	for _, e := range elements {
		s.elements[e] = struct{}{}
	}
}

// RemoveAll deletes all provided elements.
func (s *Set[T]) RemoveAll(elements ...T) {
	for _, e := range elements {
		delete(s.elements, e)
	}
}

// HasAll returns true if the set contains every provided element.
func (s *Set[T]) HasAll(elements ...T) bool {
	for _, e := range elements {
		if _, ok := s.elements[e]; !ok {
			return false
		}
	}
	return true
}

// HasAny returns true if the set contains at least one provided element.
func (s *Set[T]) HasAny(elements ...T) bool {
	for _, e := range elements {
		if _, ok := s.elements[e]; ok {
			return true
		}
	}
	return false
}

// Update mutates the set to be the union with other.
func (s *Set[T]) Update(other *Set[T]) {
	for k := range other.elements {
		s.elements[k] = struct{}{}
	}
}

// IntersectWith mutates the set to keep only elements also in other.
func (s *Set[T]) IntersectWith(other *Set[T]) {
	for k := range s.elements {
		if !other.Contains(k) {
			delete(s.elements, k)
		}
	}
}

// DifferenceWith mutates the set to remove all elements found in other.
func (s *Set[T]) DifferenceWith(other *Set[T]) {
	for k := range other.elements {
		delete(s.elements, k)
	}
}

// IsDisjoint returns true if the sets share no elements.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	smaller, larger := s, other
	if other.Size() < s.Size() {
		smaller, larger = other, s
	}
	for k := range smaller.elements {
		if larger.Contains(k) {
			return false
		}
	}
	return true
}

// Overlaps returns true if the sets share at least one element.
func (s *Set[T]) Overlaps(other *Set[T]) bool {
	return !s.IsDisjoint(other)
}

// SymmetricDifference returns elements present in exactly one of the sets.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	for k := range s.elements {
		if !other.Contains(k) {
			out.Add(k)
		}
	}
	for k := range other.elements {
		if !s.Contains(k) {
			out.Add(k)
		}
	}
	return out
}

// Pop removes and returns an arbitrary element. ok=false if empty.
func (s *Set[T]) Pop() (val T, ok bool) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, true
	}
	var zero T
	return zero, false
}

// Filter returns a new set with elements satisfying pred.
func (s *Set[T]) Filter(pred func(T) bool) *Set[T] {
	out := NewSet[T]()
	for k := range s.elements {
		if pred(k) {
			out.Add(k)
		}
	}
	return out
}

// Map returns a new set by applying f to each element.
func Map[T comparable, U comparable](in *Set[T], f func(T) U) *Set[U] {
	out := NewSet[U]()
	for k := range in.elements {
		out.Add(f(k))
	}
	return out
}

// String returns a string representation of the set.
func (s *Set[T]) String() string {
	result := "{"
	first := true
	for key := range s.elements {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v", key)
		first = false
	}
	result += "}"
	return result
}
