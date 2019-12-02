package common

import (
	"bufio"
	"os"
	"unicode"
	"unicode/utf8"
)

// ReadLines reads all lines from the given path and returns them in a slice of strings.
// If an error occurs, the function will panic.
func ReadLines(path string) []string {
	return readWithScanner(path, bufio.ScanLines)
}

// ReadCommas reads all data from the given path and returns a string slice
// containing comma-delimited parts. If a
// If an error occurs, the function will panic.
func ReadCommas(path string) []string {
	return readWithScanner(path, scanByCommas)
}

func readWithScanner(path string, splitFunc bufio.SplitFunc) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var parts []string
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		parts = append(parts, scanner.Text())
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
