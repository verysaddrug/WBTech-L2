package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Repeat(char rune, n int) []rune {
	res := make([]rune, n)
	for i := range res {
		res[i] = char
	}
	return res
}

func Parse(s string) (string, error) {
	runes := []rune(s)        // convert to runes for iterating
	result := make([]rune, 0) // result slice
	last := rune(0)           // last non-digit character
	escaped := false          // flag for escaping

	for i := 0; i < len(runes); i++ {
		if escaped {
			result = append(result, runes[i])
			last = runes[i]
			escaped = false
			continue
		}

		if runes[i] == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(runes[i]) {
			if last == 0 {
				return "", errors.New("invalid string")
			}
			num, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", errors.New("invalid string")
			}
			result = append(result, Repeat(last, num-1)...)
		} else {
			last = runes[i]
			result = append(result, last)
		}
	}

	if escaped {
		return "", errors.New("invalid string")
	}

	return string(result), nil
}

func main() {
	testCases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, tc := range testCases {
		result, err := Parse(tc)
		if err != nil {
			fmt.Printf("Error parsing %q: %v\n", tc, err)
		} else {
			fmt.Printf("Parsed %q: %q\n", tc, result)
		}
	}
}
