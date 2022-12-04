package common

import (
	"bufio"
	"io"
)

// ReadLines reads lines from a reader and calls the callback function for
// each line
func ReadLines(r io.Reader, f func(line string)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
}
