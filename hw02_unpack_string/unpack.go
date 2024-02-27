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
	} else {
		return false, 0
	}
}

func Unpack(input string) (string, error) {
	var output strings.Builder
	//var currentSymbol rune
	var previousSymbol rune
	for i, j := range input {
		if i == 0 {
			var number, _ = IsNumber(j)
			if number {
				return "", ErrInvalidString
			}
			previousSymbol = j
		} else {
			numberPrev, _ := IsNumber(previousSymbol)
			if numberPrev {
				numberCur, _ := IsNumber(j)
				if numberCur {
					return "", ErrInvalidString
				} else {
					previousSymbol = j
					// если символ последний
					if i == len([]rune((input)))-1 {
						output.WriteString(string(j))
					}
					continue
				}
			} else {
				numberCur, value := IsNumber(j)
				if numberCur {
					output.WriteString(strings.Repeat(string(previousSymbol), value))
				} else {
					output.WriteString(string(previousSymbol))
					// если символ последний
					if i == len([]rune((input)))-1 {
						output.WriteString(string(j))
					}
				}
				previousSymbol = j
			}
		}
	}

	return output.String(), nil
}
