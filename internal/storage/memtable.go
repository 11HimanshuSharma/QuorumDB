package storage

import "sync"

type MemTable struct {
	data map[string]string
	mu sync.RWMutex
}

func NewMemTable() *MemTable {
	return &MemTable{
		data: make(map[string]string),
	}
}


func (m *MemTable) Set(key, value string) {
	m.mu.Lock()

	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *MemTable) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.data[key]

	return value, exists
	
}