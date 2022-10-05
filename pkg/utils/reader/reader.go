package reader

import (
	"bufio"
	"io"
	"log"
	"os"
)

func File2Slice(path string) []string {

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	result := []string{}

	for {
		if c, sz, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(c, sz, err)
		} else {
			result = append(result, string(c))
		}
	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
