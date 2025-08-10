package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	Unpack("\\qwe\\45")
}

// ErrInvalidString - custom error for invalid input data
var ErrInvalidString error = errors.New("Invalid input string")

// Unpack - unpack string with repeats 
func Unpack(inputString string) (string, error) {
	if inputString == "" {
		return "", nil
	}

	var sb strings.Builder
	var prevSymbol rune 
	escaped := false
	hasSymbol := false


	runes := []rune(inputString)
	for i := 0; i < len(runes); i++ {
		currentSymbol := runes[i]

		if escaped { // if prev symbol is '\'
			sb.WriteRune(currentSymbol)
			prevSymbol = currentSymbol
			escaped = false
			hasSymbol = true
			continue
		}

		if currentSymbol == '\\' {
			escaped = true
		} else if unicode.IsDigit(currentSymbol) {
			if prevSymbol == 0 { // digit at first position
				return "", ErrInvalidString
			}
			count, err := strconv.Atoi(string(currentSymbol))
			if err != nil || count == 0 {
				return "", ErrInvalidString
			}
			sb.WriteString(strings.Repeat(string(prevSymbol), count-1))
		} else {
			sb.WriteRune(currentSymbol)
			prevSymbol = currentSymbol
			hasSymbol = true
		}
	}

	if escaped { // if last symbol is '\'
		return "", ErrInvalidString
	}

	if !hasSymbol { // if only digits
		return "", ErrInvalidString
	}

	return sb.String(), nil
} 