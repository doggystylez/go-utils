package safe

import (
	"sync"
	"testing"
)

func TestSafeSlice_AppendGet(t *testing.T) {
	s := NewSafeSlice[int]()

	s.Append(1)
	s.Append(2)

	if val, ok := s.Get(0); !ok || val != 1 {
		t.Errorf("Expected first element to be 1, got %v", val)
	}

	if val, ok := s.Get(1); !ok || val != 2 {
		t.Errorf("Expected second element to be 2, got %v", val)
	}
}

func TestSafeSlice_Remove(t *testing.T) {
	s := NewSafeSlice[int]()
	s.Append(1)
	s.Append(2)
	s.Append(3)

	if !s.Remove(1) {
		t.Errorf("Expected true from Remove, got false")
	}

	if val, ok := s.Get(1); !ok || val != 3 {
		t.Errorf("Expected second element to be 3 after removal, got %v", val)
	}
}

func TestSafeSlice_Size(t *testing.T) {
	s := NewSafeSlice[int]()
	s.Append(1)
	s.Append(2)

	size := s.Size()
	if size != 2 {
		t.Errorf("Expected size to be 2, got %d", size)
	}
}

func TestSafeSlice_ConcurrentAppend(t *testing.T) {
	s := NewSafeSlice[int]()
	var wg sync.WaitGroup

	// Append items concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			s.Append(val)
		}(i)
	}
	wg.Wait()

	// Check if all items were appended
	if size := s.Size(); size != 1000 {
		t.Errorf("Expected size to be 1000 after concurrent append, got %d", size)
	}
}

func TestSafeSlice_OutOfRangeGet(t *testing.T) {
	s := NewSafeSlice[int]()
	s.Append(1)
	if _, ok := s.Get(1); ok {
		t.Errorf("Expected false from Get with out-of-range index")
	}
}

func TestSafeSlice_OutOfRangeRemove(t *testing.T) {
	s := NewSafeSlice[int]()
	s.Append(1)
	if s.Remove(1) {
		t.Errorf("Expected false from Remove with out-of-range index")
	}

	if s.Remove(-1) {
		t.Errorf("Expected false from Remove with negative index")
	}
}
