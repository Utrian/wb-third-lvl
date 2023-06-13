package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Реализовать утилиту фильтрации - аналог grep

var (
	fAfter      int
	fBefore     int
	fContext    int
	fCount      bool
	fIgnoreCase bool
	fInvert     bool
	fFixed      bool
	fLineNum    bool
)

func initFlags() {
	flag.IntVar(&fAfter, "A", 0, "Print  NUM  lines of trailing context after matching lines.")
	flag.IntVar(&fBefore, "B", 0, "Print  NUM  lines of leading context before matching lines.")
	flag.IntVar(&fContext, "C", 0, "Print   NUM   lines  of  output  context.")
	flag.BoolVar(&fCount, "c", false, "Suppress normal output; instead print a count of matching lines for each input file.")
	flag.BoolVar(&fIgnoreCase, "i", false, "Ignore case distinctions in patterns and input data.")
	flag.BoolVar(&fInvert, "v", false, "Invert the sense of matching, to select non-matching lines.")
	flag.BoolVar(&fFixed, "F", false, "Interpret PATTERNS as fixed strings, not regular expressions.")
	flag.BoolVar(&fLineNum, "n", false, "Prefix each line of output with the 1-based line number within its input file.")

	flag.Parse()
}

func main() {
	initFlags()

	path := "develop/dev05/text.txt"
	file, _ := os.Open(path)

	matched := grep(file, []rune("soon"))
	if fCount {
		fmt.Println(len(matched))
	} else {
		for _, v := range matched {
			fmt.Println(string(v))
		}
	}
}

func grep(file *os.File, search []rune) [][]rune {
	var result [][]rune

	sc := bufio.NewScanner(file)

	var match int
	// var countLines int

	// Читаем файл
Scan:
	for sc.Scan() {
		var lineNumber int

		// Получаем строки и преобразуем в руны
		text := sc.Text()
		if fIgnoreCase {
			text = strings.ToLower(string(text))
		}
		line := []rune(text)
		// Посимвольно проверяем совпадения с шаблоном
		for i, letter := range line {
			if letter == search[match] {
				// Проверка на то, что шаблон закончился
				if match == len(search)-1 {
					// скорее всего флаги должны быть здесь
					firstLetter := i - match
					lastLetter := i + 1

					switch {
					case fAfter > 0:
						lastLetter = i + 1 + fAfter
						max := len(line)
						if lastLetter > max {
							lastLetter = max
						}
					case fBefore > 0:
						firstLetter -= fBefore
						if firstLetter < 0 {
							firstLetter = 0
						}
					case fContext > 0:
						firstLetter -= fContext
						if firstLetter < 0 {
							firstLetter = 0
						}

						lastLetter = i + 1 + fContext
						max := len(line)
						if lastLetter > max {
							lastLetter = max
						}
					case fFixed:
						preFirstLetter := firstLetter - 1
						postLastLetter := lastLetter + 1
						if preFirstLetter > 0 && line[preFirstLetter] != ' ' {
							continue
						}
						if postLastLetter < len(line) && line[postLastLetter] != ' ' {
							continue
						}
					}

					match = 0 // после успешного нахождения сбрасываем счетчик

					result = append(result, line[firstLetter:lastLetter])

					if fCount {
						continue Scan
					}

				} else {
					match++
				}

			} else {
				match = 0 // несовпадение шаблона - сбрасываем счетчик
			}
		}
		lineNumber++
	}
	return result
}
