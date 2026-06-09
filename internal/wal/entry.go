package wal


import (
	"fmt"
	"time"
)


type Entry struct {
	TimeStamp int64
	Operation string
	Key  string
	Value string
}


func NewEntry(op, key, value string) Entry {
	return Entry{
		TimeStamp: time.Now().UnixNano(),
		Operation: op,
		Key: key,
		Value: value,
	}
}

func (e Entry) Serialize() string {
	return fmt.Sprintf(
		"%d|%s|%s|%s",
		e.TimeStamp,
		e.Operation,
		e.Key,
		e.Value,
	)
}