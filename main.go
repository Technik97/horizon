package main

import (
	"fmt"
	"log"

	"github.com/ergochat/readline"
)

func main() {
	fmt.Println("Horizon version 0.0.1")
	fmt.Println("Press Ctrl+c to Exit")

	rl, err := readline.New("hrzn> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.ReadLine()
		if err != nil {
			break // io.EOF or readline.ErrInterrupt (Ctrl+C)
		}

		fmt.Printf("%s\n", line)
	}
}
