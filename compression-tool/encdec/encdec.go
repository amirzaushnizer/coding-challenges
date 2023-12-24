package encdec

import (
	"fmt"
	"log"
	"os"
)

func WriteEncodedFile(inputFile *os.File,
	encodingTable map[byte]string,
	endcodedFilePath string) {

	// create encoded file
	encodedFile, err := os.Create(endcodedFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer encodedFile.Close()
	// write the encoding to the file
	for char, path := range encodingTable {
		//write beginning of new recored as |
		encodedFile.WriteString("|")
		encodedFile.WriteString(path)
		// write delimiter for end of key as :
		encodedFile.WriteString(":")
		encodedFile.WriteString(string(char))
	}

	// write delimiter for end of encoding as _
	encodedFile.WriteString("_")

	// write the encoded file as bits
	readBuff := make([]byte, 1)
	bytePack := byte(0)
	bytePackPos := 0

	for {
		_, err := inputFile.Read(readBuff)
		if err != nil {
			break
		}
		bitsString := encodingTable[readBuff[0]]
		for _, bit := range bitsString {
			if bit == '1' {
				bytePack |= 1 << bytePackPos
			}
			bytePackPos++
			if bytePackPos == 8 {
				encodedFile.Write([]byte{bytePack})
				bytePackPos = 0
				bytePack = byte(0)
			}

		}
	}

	// write the last byte and the number of bits used in the last byte
	if bytePackPos > 0 {
		encodedFile.Write([]byte{bytePack})
		encodedFile.Write([]byte{byte(bytePackPos)})
	}
}

func ReadEncodingTable(encodedFile *os.File) map[string]byte {
	encodingTable := make(map[string]byte)

	// read the encoding table
	readBuff := make([]byte, 1)
	encoding := ""
	for {
		_, err := encodedFile.Read(readBuff)
		if err != nil {
			break
		}
		if readBuff[0] == '|' {
			encoding = ""
		} else if readBuff[0] == ':' {
			encodedFile.Read(readBuff)
			encodingTable[encoding] = readBuff[0]
		} else if readBuff[0] == '_' {
			break
		} else {
			encoding += string(readBuff[0])
		}
	}

	return encodingTable
}

func DecodeFile(encodedFile *os.File, encodingTable map[string]byte, outputFilePath string) {
	//print the current offset of the encoded file
	fmt.Println(encodedFile.Seek(0, 1))
	// create output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// read the encoded file as bits
	readBuff := make([]byte, 1)
	bitString := ""

	for {
		_, err := encodedFile.Read(readBuff)
		if err != nil {
			break
		}
		for i := 0; i < 8; i++ {
			bit := (readBuff[0] >> i) & 1
			if bit == 1 {
				bitString += "1"
			} else {
				bitString += "0"
			}
			if char, ok := encodingTable[bitString]; ok {
				outputFile.Write([]byte{char})
				bitString = ""
			}
		}
	}
}
