package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No text file provided as argument")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error while reading file: ", err)
	}

	io.Copy(os.Stdout, file)
}
