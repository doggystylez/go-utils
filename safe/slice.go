package safe

import (
	"sync"
)

// SafeSlice is a thread-safe slice.
type SafeSlice[T any] struct {
	mu    sync.RWMutex
	items []T
}

// NewSafeSlice creates a new SafeSlice.
func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{items: make([]T, 0)}
}

// Append appends the given item to the slice.
func (s *SafeSlice[T]) Append(item T) {
	s.mu.Lock()
	s.items = append(s.items, item)
	s.mu.Unlock()
}

// Remove removes the item at the given index from the slice.
func (s *SafeSlice[T]) Remove(index int) bool {
	s.mu.Lock()
	if index < 0 || index >= len(s.items) {
		s.mu.Unlock()
		return false
	}
	s.items = append(s.items[:index], s.items[index+1:]...)
	s.mu.Unlock()
	return true
}

// Get returns the item at the given index from the slice.
func (s *SafeSlice[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var zero T // Create a zero value for type T
	if index < 0 || index >= len(s.items) {
		return zero, false
	}
	return s.items[index], true
}

// Size returns the number of items in the slice.
func (s *SafeSlice[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Items returns a copy of the slice.
func (s *SafeSlice[T]) Items() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	itemsCopy := make([]T, len(s.items))
	copy(itemsCopy, s.items)
	return itemsCopy
}
