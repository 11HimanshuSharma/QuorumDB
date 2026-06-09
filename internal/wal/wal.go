package wal


import (
	"bufio"
	"os"
	"sync"
)

type WAL struct {
	file *os.File
	writer *bufio.Writer
	mu sync.Mutex
}


func NewWal(path string) (*WAL, error) {
	file, err := os.OpenFile(
		path,
		os.O_CREATE | os.O_WRONLY | os.O_APPEND,
		0644,
	)

	return &WAL{
		file: file,
		writer: bufio.NewWriter(file),
	}, err
}


func (w *WAL) Append(entry Entry) error {
	w.mu.Lock()

	defer w.mu.Unlock()

	serialized := entry.Serialize()

	if err != nil {
		return err
	}
	// flush buffer -> OS buffer
	err = w.writer.Flush()

	if err != nil {
		return err
	}

	// OS buffer -> dist
	// durability guarantee
	err = w.file.Sync()

	if err != nil {
		return err
	}

	return nil


}

func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()


	if err != nil {
		return err
	}

	err = w.file.Sync()

	if err != nil {
		return err
	}
	return w.file.Close()
}