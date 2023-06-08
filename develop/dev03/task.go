package main

// Отсортировать строки
// На вход - файл, на выход - файл с отсортированными строками
// Реализовать флаги:
// -k — указание колонки для сортировки (слова в строке могут
// выступать в качестве колонок, по умолчанию разделитель —
// пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// Основные флаги
	fColumn  int
	fNumeric bool
	fReverse bool
	fUnique  bool

	// Дополнительные флаги
	fMonth    bool
	fBlanks   bool
	fCheck    bool
	fHumanNum bool
)

// Инициализация флагов
func initFlags() {
	flag.IntVar(&fColumn, "k", 0, "column number for sorting")
	flag.BoolVar(&fNumeric, "n", false, "compare according to string numerical value")
	flag.BoolVar(&fReverse, "r", false, "reverse the result of comparisons")
	flag.BoolVar(&fUnique, "u", false, "output only the first of an equal run")

	flag.BoolVar(&fMonth, "M", false, "compare (unknown) < 'JAN' < ... < 'DEC'")
	flag.BoolVar(&fBlanks, "b", false, "ignore leading blanks")
	flag.BoolVar(&fCheck, "c", false, "check for sorted input; do not sort")
	flag.BoolVar(&fHumanNum, "h", false, "compare human readable numbers (e.g., 2K 1G)")

	flag.Parse()
}

func main() {
	initFlags()

	path := "/home/paul/go/src/github.com/wb-third-lvl/develop/dev03/text.txt"

	// Открываем файл
	file, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
	}

	// Получим отсортированный файл
	sortedFile, err := sortFile(file)
	if err != nil {
		logrus.Error(err)
	}

	sortedFile.Close()
	file.Close()
}

func sortFile(file *os.File) (*os.File, error) {
	// Формируем слайс строк из строк файла
	lines, err := scanFile(file)
	if err != nil {
		return nil, err
	}

	// Сортировка слайса строк
	lines = sortLines(lines)

	// Создаем файл для отсортированных строк
	resultFile, err := os.Create(createResultFileName(file.Name()))
	if err != nil {
		return nil, err
	}

	// Следующие два блока почти одинаковые
	// первый реализован для флага 'u'
	// Второй для кейса без флага 'u'
	// Можно было бы проверять флаг внутри for,
	// но тогда бы кол-во проверок было бы
	// равно кол-ву строк. Поэтому я решил, что
	// оптимальнее повторить код.
	if fUnique {
		var last string

		for _, line := range lines {
			if line != last {
				fmt.Fprintln(resultFile, line)
				last = line
			}
		}
		return resultFile, nil
	}

	for _, line := range lines {
		fmt.Fprintln(resultFile, line)
	}

	return resultFile, nil
}

func scanFile(file *os.File) ([]string, error) {
	var lines []string

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

// Сортировка строк слайса, тут основная часть реализаций флагов
func sortLines(lines []string) []string {
	sort.Slice(lines, func(i int, j int) bool {
		a := strings.Fields(lines[i])
		b := strings.Fields(lines[j])

		// Если установлен флаг 'k'
		// Не попадаем в этот блок если колонок внутри строки меньше
		// чем значение флага, т.к. иначе будет выход за range
		if fColumn > 0 && fColumn <= len(a) && fColumn <= len(b) {
			// Берем срезы чтобы выбранная колонка/слово было первым
			a = a[fColumn-1 : fColumn]
			b = b[fColumn-1 : fColumn]
		}

		if fNumeric {
			num1, err := strconv.Atoi(a[0])
			if err != nil {
				// В случае, если строка не представлена числом
				// используем обычное сравнение. В итоге сначала
				// будут отсортированные строки по флагу 'n', затем
				// отсортированные строки по стандартному сценарию
				return strings.Compare(a[0], b[0]) < 0
			}
			num2, err := strconv.Atoi(b[0])
			if err != nil {
				return strings.Compare(a[0], b[0]) < 0
			}

			return num1 < num2
		}

		less := strings.Compare(a[0], b[0]) < 0
		if fReverse {
			less = !less
		}

		return less
	})

	return lines
}

// На основании пути оригинально файла, формирует
// путь для нового файла в ту же директорию
// имя нового файла "sorted_[oldFileName]"
func createResultFileName(path string) string {
	pathSlice := strings.FieldsFunc(
		path,
		func(c rune) bool {
			return c == '/'
		},
	)
	newFileName := "sorted_" + pathSlice[len(pathSlice)-1]
	pathSlice[len(pathSlice)-1] = newFileName

	newPath := strings.Join(pathSlice, "/")

	return "/" + newPath
}
