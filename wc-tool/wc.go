package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func countBytes(file *os.File) int {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(fileInfo.Size())
}

func countLines(file *os.File) int {
	lineCount := 0
	var buf = make([]byte, 1)
	for {
		_, err := file.Read(buf)
		if err != nil {
			break
		}
		if buf[0] == '\n' {
			lineCount++
		}
	}

	file.Seek(0, 0)
	return lineCount
}

func countWords(file *os.File) int {
	wordCount := 0
	inWord := false

	buf := make([]byte, 1)

	var isSpace bool

	for {
		_, err := file.Read(buf)
		if err != nil {
			break
		}

		isSpace = unicode.IsSpace(rune(buf[0]))

		if isSpace && inWord {
			wordCount++
			inWord = false
		} else if !isSpace {
			inWord = true
		}
	}
	file.Seek(0, 0)
	return wordCount
}

func countChars(file *os.File) int {
	charCount := 0

	buf := make([]byte, 4096)

	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		charCount += utf8.RuneCount(buf[:n])
	}
	// reset file pointer to the beginning of the file
	file.Seek(0, 0)
	return charCount
}

func main() {
	filePath := os.Args[len(os.Args)-1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cFlag := flag.Bool("c", false, "print the byte counts")
	lFlag := flag.Bool("l", false, "print the newline counts")
	wFlag := flag.Bool("w", false, "print the word counts")
	mFlag := flag.Bool("m", false, "print the character counts")

	flag.Parse()
	// if no flag is provided, print all counts except character count
	if !*cFlag && !*lFlag && !*wFlag && !*mFlag {
		// print all counts in a single line separated by tabs
		fmt.Printf("%d\t%d\t%d\n", countLines(file), countWords(file), countBytes(file))
	}

	if *cFlag {
		fmt.Println(countBytes(file))
	}

	if *lFlag {
		fmt.Println(countLines(file))
	}

	if *wFlag {
		fmt.Println(countWords(file))
	}

	if *mFlag {
		fmt.Println(countChars(file))
	}
}
