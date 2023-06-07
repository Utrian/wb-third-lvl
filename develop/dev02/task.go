package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Создать Go-функцию, осуществляющую примитивную распаковку
// строки, содержащую повторяющиеся символы/руны, например:
// ● "a4bc2d5e" => "aaaabccddddde"
// ● "abcd" => "abcd"
// ● "45" => "" (некорректная строка)
// ● "" => ""

// Реализовать поддержку escape-последовательностей.
// Например:
// ● qwe\4\5 => qwe45 (*)
// ● qwe\45 => qwe44444 (*)
// ● qwe\\5 => qwe\\\\\ (*)

// В случае если была передана некорректная строка, функция
// должна возвращать ошибку. Написать unit-тесты.

func main() {
	s, err := UnpackStr(``)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

func UnpackStr(s string) (string, error) {
	var last rune
	var sb strings.Builder

	for _, r := range s {
		if unicode.IsDigit(r) {
			// Проверка и обработка некорректных строк:
			// Если два последних символа были числами и
			// если число идет первым в строке
			if unicode.IsDigit(last) || last == 0 {
				err := fmt.Sprintf("the string '%s' isn't correct", s)
				return s, errors.New(err)
			}

			// Получаем число повторений
			// И если число повторений не отрицательное, то
			// добавляем в билдер
			if repeat := int(r - '1'); repeat >= 0 {
				str := strings.Repeat(string(last), repeat)
				sb.WriteString(str)
			} else {
				// Обработка случая, когда опльзователь ввел 0
				// иначе будет паника вызванная strings.Repeat
				return s, errors.New("don`t use '0' in the string")
			}
		} else {
			sb.WriteRune(r)
		}
		last = r
	}

	return sb.String(), nil
}
