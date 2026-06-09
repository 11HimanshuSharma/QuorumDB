package storage


import (
	"bufio",
	"os",
	"strings"
)


func (e *Engine) Recover(walPath string) error {
	file, err := os.Open(walPath)

	if err != nil {
		return err
	}

	defer file.Close()
	

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "|")

		if len(parts) != 4 {
			continue
		}

		operation := parts[1]
		key := parts[2]
		value := parts[3]

		switch operation {

		case "SET":
			e.memTable.Set(key, value)
		}
	}
	return scanner.Err()
}