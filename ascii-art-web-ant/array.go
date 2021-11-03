package ascii

import (
	"bufio"
	"io"
	"os"
)

func Array(file *os.File) []string {

	var lines []string
	scanner := bufio.NewScanner(io.Reader(file))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
