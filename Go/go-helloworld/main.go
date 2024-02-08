package main

import (
	"fmt"
	"os"
)

func main() {
	//var args []string  // declaring string
	args := os.Args

	if len(args) < 1 {
		fmt.Printf("Usage: ./hello-world <argument>\n")
		os.Exit(1)
	}

	fmt.Printf("Hello World!\nArguments: %v\n", args[1]) // f stands for format
}
