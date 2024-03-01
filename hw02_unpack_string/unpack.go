package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func IsNumber(input rune) (bool, int) {
	result, err := strconv.Atoi(string(input))
	if err == nil {
		return true, result
	}
	return false, 0
}

func Unpack(input string) (string, error) {
	var output strings.Builder
	var previousSymbol rune
	for i, currentSymbol := range input {
		if i == 0 {
			number, _ := IsNumber(currentSymbol)
			if number {
				return "", ErrInvalidString
			}
			previousSymbol = currentSymbol
			continue
		}
		numberPrev, _ := IsNumber(previousSymbol)
		numberCur, valueCur := IsNumber(currentSymbol)
		switch {
		case numberPrev && numberCur:
			return "", ErrInvalidString
		case numberPrev && !numberCur:
			previousSymbol = currentSymbol
			if i == len([]rune((input)))-1 {
				output.WriteString(string(currentSymbol))
			}
			continue
		case !numberPrev && numberCur:
			output.WriteString(strings.Repeat(string(previousSymbol), valueCur))
		case !numberPrev && !numberCur:
			output.WriteString(string(previousSymbol))
			if i == len([]rune((input)))-1 {
				output.WriteString(string(currentSymbol))
			}
		}
		previousSymbol = currentSymbol
	}

	return output.String(), nil
}
