package file

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ToSlice(path string) []float64 {
	file, err := os.Open(path)
	if err != nil {
		panic("foo")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	result := []float64{}

	for {
		if c, sz, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(c, sz, err)
		} else {
			result = append(result, float64(c))
		}
	}

	return result
}
