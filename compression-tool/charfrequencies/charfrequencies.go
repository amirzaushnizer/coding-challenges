package charfrequencies

import "os"

func GetCharFrequencies(file *os.File) map[byte]int {
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
