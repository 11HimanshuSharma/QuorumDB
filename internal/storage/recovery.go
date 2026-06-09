package storage


import (
	"bufio",
	"os",
	"strings"
)


func (e *Engine) Recover(walPath string) error {
	file, err := os.Open(walPath)

	if err !
}