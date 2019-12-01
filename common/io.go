package common

import (
	"bufio"
	"os"
)

// ReadLines reads all lines from the given path and returns them in a slice of strings.
// If an error occurs, the function will panic.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}
