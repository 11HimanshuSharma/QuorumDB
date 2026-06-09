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

	}
}