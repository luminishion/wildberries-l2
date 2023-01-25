package main

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

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) string {
	chars := []rune(str)

	var ret strings.Builder
	escNext := false
	var lastChar rune = 0

	for i := 0; i < len(chars); i++ {
		if !escNext && chars[i] == rune('\\') {
			escNext = true
			continue
		}

		if escNext || !unicode.IsDigit(chars[i]) {
			escNext = false

			if lastChar != 0 {
				ret.WriteRune(lastChar)
			}

			lastChar = chars[i]

			continue
		}

		if lastChar == 0 {
			return ""
		}

		numStart := i
		numEnd := i + 1
		for numEnd < len(chars) && unicode.IsDigit(chars[numEnd]) {
			numEnd++
		}

		numStr := string(chars[numStart:numEnd])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return ""
		}

		ret.WriteString(strings.Repeat(string(lastChar), num))
		lastChar = 0

		i = numEnd - 1
	}

	if lastChar != 0 {
		ret.WriteRune(lastChar)
	}

	return ret.String()
}

func main() {
	fmt.Println(Unpack("a4bc2d5e") == "aaaabccddddde")
	fmt.Println(Unpack("abcd") == "abcd")
	fmt.Println(Unpack("45") == "")
	fmt.Println(Unpack("") == "")
	fmt.Println(Unpack(`qwe\4\5`) == `qwe45`)
	fmt.Println(Unpack(`qwe\45`) == `qwe44444`)
	fmt.Println(Unpack(`qwe\\5`) == `qwe\\\\\`)
}
