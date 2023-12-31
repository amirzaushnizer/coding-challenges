package main

import (
	"compression/charfrequencies"
	"compression/encdec"
	"compression/hufftree"
	"flag"
	"log"
	"os"
)

func main() {

	decodingFlag := flag.Bool("d", false, "decode the file")
	flag.Parse()

	if !*decodingFlag {
		if len(os.Args) < 3 {
			log.Fatal("Please provide a file path and an output file path")
			return
		}
		filePath := os.Args[1]
		endcodedFilePath := os.Args[2]

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		frequencies := charfrequencies.GetCharFrequencies(file)

		huffHeap := hufftree.CreateHuffHeap(frequencies)
		huffTree := huffHeap.ToHuffTree()

		encodingTable := huffTree.Encode()
		encdec.WriteEncodedFile(file, encodingTable, endcodedFilePath)
	} else {
		if len(os.Args) < 4 {
			log.Fatal("Please provide a file path and an output file path")
			return
		}
		endcodedFilePath := os.Args[2]
		outputFilePath := os.Args[3]

		file, err := os.Open(endcodedFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		encodingTable := encdec.ReadEncodingTable(file)
		encdec.DecodeFile(file, encodingTable, outputFilePath)
	}
}
