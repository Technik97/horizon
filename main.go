package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Horizon version 0.0.1")
	fmt.Println("Press Ctrl+c to Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("hrzn>")

		if !scanner.Scan() {
			break // EOF or error
		}

		input := strings.TrimSpace(scanner.Text())

		// exit commands
		if input == "exit" || input == "quit" {
			break
		}

		// Skip empty input
		if input == "" {
			continue
		}

		fmt.Printf("%q\n", input)
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		}
	}

	fmt.Println("Goodbye!")
}
