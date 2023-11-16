package safe

import (
	"sync"
)

// UniqueSafeSlice is a thread-safe slice that only allows unique items.
type UniqueSafeSlice[T comparable] struct {
	mu     sync.RWMutex
	items  []T
	unique map[T]struct{}
}

// NewUniqueSafeSlice creates a new UniqueSafeSlice.
func NewUniqueSafeSlice[T comparable]() *UniqueSafeSlice[T] {
	return &UniqueSafeSlice[T]{
		items:  make([]T, 0),
		unique: make(map[T]struct{}),
	}
}

// Append appends the given item to the slice if it does not already exist in the slice.
func (s *UniqueSafeSlice[T]) Append(item T) {
	s.mu.Lock()
	if _, exists := s.unique[item]; !exists {
		s.items = append(s.items, item)
		s.unique[item] = struct{}{}
	}
	s.mu.Unlock()
}

// Remove removes the item at the given index from the slice.
func (s *UniqueSafeSlice[T]) Remove(index int) bool {
	s.mu.Lock()
	if index < 0 || index >= len(s.items) {
		s.mu.Unlock()
		return false
	}
	removedItem := s.items[index]
	delete(s.unique, removedItem)
	s.items = append(s.items[:index], s.items[index+1:]...)
	s.mu.Unlock()
	return true
}

// Get returns the item at the given index from the slice.
func (s *UniqueSafeSlice[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var zero T // Create a zero value for type T
	if index < 0 || index >= len(s.items) {
		return zero, false
	}
	return s.items[index], true
}

// Size returns the number of items in the slice.
func (s *UniqueSafeSlice[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Items returns a copy of the slice.
func (s *UniqueSafeSlice[T]) Items() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	itemsCopy := make([]T, len(s.items))
	copy(itemsCopy, s.items)
	return itemsCopy
}
