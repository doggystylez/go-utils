package safe

import (
	"sync"
)

// OrderedMap is a thread-safe map that preserves the insertion order of keys.
type OrderedMap[K comparable, V any] struct {
	sync.RWMutex
	keys   []K
	values map[K]V
}

// NewOrderedMap creates a new OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

// Set adds a new key-value pair to the map or updates the value of an existing key.
func (om *OrderedMap[K, V]) Set(key K, value V) {
	om.Lock()
	defer om.Unlock()
	if _, exists := om.values[key]; !exists {
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	om.RLock()
	defer om.RUnlock()
	value, exists := om.values[key]
	if !exists {
		var zero V // The zero value for the type V
		return zero, false
	}
	return value, true
}

// Delete removes the key-value pair associated with the key.
func (om *OrderedMap[K, V]) Delete(key K) {
	om.Lock()
	defer om.Unlock()
	if _, exists := om.values[key]; exists {
		delete(om.values, key)
		// Remove the key from the keys slice
		for i, k := range om.keys {
			if k == key {
				om.keys = append(om.keys[:i], om.keys[i+1:]...)
				break
			}
		}
	}
}

// Keys returns a slice of keys in the order they were added.
func (om *OrderedMap[K, V]) Keys() []K {
	om.RLock()
	defer om.RUnlock()
	keysCopy := make([]K, len(om.keys))
	copy(keysCopy, om.keys)
	return keysCopy
}

// Values returns a slice of values in the order their corresponding keys were added.
func (om *OrderedMap[K, V]) Values() []V {
	om.RLock()
	defer om.RUnlock()
	values := make([]V, len(om.keys))
	for i, key := range om.keys {
		values[i] = om.values[key]
	}
	return values
}

// Len returns the number of key-value pairs in the map.
func (om *OrderedMap[K, V]) Len() int {
	om.RLock()
	defer om.RUnlock()
	return len(om.keys)
}
