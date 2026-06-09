package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"quorumdb/internal/storage"
)

func main() {
	engine, err := storage.NewEngine("data/wal/wal.log")

	if err != nil {
		panic(err)
	}

	defer engine.Close()

	// Recover
	err = engine.Recover("data/wal/wal.log")

	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to QuorumDB!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")

		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])
		switch command {
		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET <key> <value>")
				continue
			}

			err := engine.Set(parts[1], parts[2])

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("OK")

		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET <key>")
				continue
			}
			value, exists := engine.Get(parts[1])

			if !exists {
				fmt.Println("key not found")
				continue
			} 

			fmt.Println(value)

		default:
			fmt.Println("Unknown command")
		}

	}
}
