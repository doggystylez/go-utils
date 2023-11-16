package safe

import (
	"sync"
	"testing"
)

func TestUniqueSafeSlice_Append(t *testing.T) {
	s := NewUniqueSafeSlice[int]()

	s.Append(1)
	s.Append(1) // This should not be appended

	if s.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", s.Size())
	}
}

func TestUniqueSafeSlice_Remove(t *testing.T) {
	s := NewUniqueSafeSlice[int]()
	s.Append(1)
	s.Append(2)
	s.Append(3)

	if !s.Remove(1) {
		t.Errorf("Expected true from Remove, got false")
	}

	if s.Size() != 2 {
		t.Errorf("Expected size to be 2 after removal, got %d", s.Size())
	}
}

func TestUniqueSafeSlice_ConcurrentAppend(t *testing.T) {
	s := NewUniqueSafeSlice[int]()
	var wg sync.WaitGroup

	// Attempt to append the same item concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Append(1)
		}()
	}
	wg.Wait()

	// Check if only one item was appended
	if size := s.Size(); size != 1 {
		t.Errorf("Expected size to be 1 after concurrent append of the same item, got %d", size)
	}
}

func TestUniqueSafeSlice_AppendDifferent(t *testing.T) {
	s := NewUniqueSafeSlice[int]()
	s.Append(1)
	s.Append(2) // Different value should be appended

	if size := s.Size(); size != 2 {
		t.Errorf("Expected size to be 2 after appending a different item, got %d", size)
	}
}
