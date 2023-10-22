package safe

import (
	"sync"
)

type SafeSlice struct {
	sync.RWMutex
	items []interface{}
}

// NewSafeSlice creates a new SafeSlice.
func NewSafeSlice() *SafeSlice {
	return &SafeSlice{
		items: make([]interface{}, 0),
	}
}

// Append adds an item to the slice.
func (s *SafeSlice) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.items = append(s.items, item)
}

// Get returns the item at the given index.
func (s *SafeSlice) Get(index int) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()
	if index < 0 || index >= len(s.items) {
		return nil, false
	}
	return s.items[index], true
}

// Remove removes the item at the given index.
func (s *SafeSlice) Remove(index int) bool {
	s.Lock()
	defer s.Unlock()
	if index < 0 || index >= len(s.items) {
		return false
	}
	s.items = append(s.items[:index], s.items[index+1:]...)
	return true
}

// Size returns the number of elements in the slice.
func (s *SafeSlice) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.items)
}
