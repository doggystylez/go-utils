package safe

import (
	"sync"
	"testing"
)

func TestOrderedMap_SetGet(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("one", 1)
	om.Set("two", 2)

	if val, ok := om.Get("one"); !ok || val != 1 {
		t.Errorf("Expected value for key 'one' to be 1, got %v", val)
	}

	if val, ok := om.Get("two"); !ok || val != 2 {
		t.Errorf("Expected value for key 'two' to be 2, got %v", val)
	}
}

func TestOrderedMap_Delete(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("one", 1)
	om.Set("two", 2)
	om.Set("three", 3)

	om.Delete("two")

	if _, ok := om.Get("two"); ok {
		t.Errorf("Expected key 'two' to be deleted")
	}

	if om.Len() != 2 {
		t.Errorf("Expected length to be 2 after deletion, got %d", om.Len())
	}
}

func TestOrderedMap_Order(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("one", 1)
	om.Set("two", 2)
	om.Set("three", 3)

	keys := om.Keys()
	expectedKeys := []string{"one", "two", "three"}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Errorf("Expected key order %v, got %v", expectedKeys, keys)
			break
		}
	}
}

func TestOrderedMap_ConcurrentSet(t *testing.T) {
	om := NewOrderedMap[int, int]()
	var wg sync.WaitGroup

	// Set key-value pairs concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			om.Set(key, key*10)
		}(i)
	}
	wg.Wait()

	// Check if all key-value pairs were set
	if length := om.Len(); length != 1000 {
		t.Errorf("Expected length to be 1000 after concurrent set, got %d", length)
	}
}

func TestOrderedMap_ConcurrentGetSet(t *testing.T) {
	om := NewOrderedMap[int, int]()
	var wg sync.WaitGroup

	// Set initial key-value pairs
	for i := 0; i < 100; i++ {
		om.Set(i, i*10)
	}

	// Concurrently set key-value pairs
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			om.Set(key, key*20)
		}(i)
	}

	// Concurrently get key-value pairs
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			om.RLock() // Use RLock to simulate the Get method's locking behavior
			val, ok := om.values[key]
			om.RUnlock()
			if !ok {
				t.Errorf("Expected key %d to exist", key)
			}
			if val != key*10 && val != key*20 {
				t.Errorf("Expected value for key %d to be %d or %d, got %d", key, key*10, key*20, val)
			}
		}(i)
	}
	wg.Wait()
}
func TestOrderedMap_OrderAfterDelete(t *testing.T) {
	om := NewOrderedMap[int, int]()

	// Set key-value pairs
	for i := 0; i < 5; i++ {
		om.Set(i, i*10)
	}

	// Delete a key
	om.Delete(2)

	// Check the order of remaining keys
	expectedKeys := []int{0, 1, 3, 4}
	keys := om.Keys()
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Errorf("Expected key order %v, got %v", expectedKeys, keys)
			break
		}
	}
}
