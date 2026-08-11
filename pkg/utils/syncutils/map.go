package syncutils

import (
	"maps"
	"slices"
	"sync"
)

// Map stores the examples and also make it memory safe.
type Map struct {
	sync.RWMutex
	data map[string]string
}

// NewMap creates new Map
func NewMap() *Map {
	return &Map{
		data: make(map[string]string),
	}
}

// SetKV sets a key-value pair
func (m *Map) SetKV(key string, value string) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

// GetValue gets the value of a key-value pair
func (m *Map) GetValue(key string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

// GetAllKeys gets all the key of the map
func (m *Map) GetAllKeys() []string {
	m.RLock()
	defer m.RUnlock()
	result := slices.Collect(maps.Keys(m.data))
	return result
}
