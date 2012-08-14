package main

import (
	"fmt"
	"github.com/joshlf13/grammar"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <grammar file>\n", os.Args[0])
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(2)
	}

	g, err := grammar.New(file)

	if err != nil {
		fmt.Printf("Error creating grammar: \n%v\n", err)
		os.Exit(3)
	}

	fmt.Println("Press enter to speak")
	for {
		fmt.Scanln()
		if err := g.Speak(os.Stdout); err != nil {
			fmt.Printf("Error writing to stdout: %v\n", err)
			os.Exit(4)
		}
		if _, err := os.Stdout.Write([]byte{'.'}); err != nil {
			fmt.Printf("Error writing to stdout: %v\n", err)
			os.Exit(4)
		}
	}
}
