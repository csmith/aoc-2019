package common

import (
	"bufio"
	"os"
	"unicode"
	"unicode/utf8"
)

// ReadFileAsInts reads all lines from the given path and returns them in a slice of ints.
// If an error occurs, the function will panic.
func ReadFileAsInts(path string) []int {
	return readIntsWithScanner(path, bufio.ScanLines)
}

// ReadCsvAsInts reads all data from the given path and returns an int slice
// containing comma-delimited parts.  If an error occurs, the function will panic.
func ReadCsvAsInts(path string) []int {
	return readIntsWithScanner(path, scanByCommas)
}

func readIntsWithScanner(path string, splitFunc bufio.SplitFunc) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var parts []int
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		parts = append(parts, MustAtoi(scanner.Text()))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return parts
}

// scanByCommas is a split function for a Scanner that returns each
// comma-separated section of text, with surrounding spaces deleted.
// The definition of space is set by unicode.IsSpace.
func scanByCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	// Scan until comma, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' || unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
