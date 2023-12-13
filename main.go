package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := parseJSON(filename)
	if err != nil {
		fmt.Printf("error parsing JSON from file %s: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
