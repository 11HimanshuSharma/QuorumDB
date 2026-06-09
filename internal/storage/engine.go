package storage

import (
	"quarkdb/internal/wal"
)

type Engine struct {
	memTable *MemTable
	wal *wal.WAL
}

func NewEngine(walPath string) (*Engine, error) {

	walInstance, err := wal.NewWal(walPath)

	if err != nil {
		return nil, err
	}

	engine := &Engine {
		memTable: NewMemTable(), 
		wal: walInstance,
	}

	return engine, nil
}

func (e *Engine) Set(key, value string) error {

	entry := wal.NewEntry(
		"SET",
		key, 
		value,
	)

	// write ahead log first
	// durability before memory update
	err := e.wal.Append(entry)
	
	if err != nil {
		return err
	}

	// update memory after WAL success

	e.memTable.Set(key, value)
	return nil
}


func (e *Engine) Get(key string) (string, bool) {
	return e.memTable.Get(key)
}
func (e *Engine) Close() error {
	return e.wal.Close()
}