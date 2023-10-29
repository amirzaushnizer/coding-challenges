package main

import (
	"fmt"
	"log"
	"os"
)

func getCharFrequencies(file *os.File) map[byte]int {
	frequencies := make(map[byte]int)

	buf := make([]byte, 1)
	for {
		_, err := file.Read(buf)
		if err != nil {
			break
		}
		frequencies[buf[0]]++
	}
	file.Seek(0, 0)
	return frequencies
}

func printCharFrequencies(frequencies map[byte]int) {
	for char, count := range frequencies {
		fmt.Printf("%q: %d\n", char, count)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// printCharFrequencies(getCharFrequencies(file))
	frequencies := getCharFrequencies(file)

}
