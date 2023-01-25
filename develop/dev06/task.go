package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	fields    string
	delimiter string
	separated bool
}

func parseFlags() Flags {
	var f Flags

	flag.StringVar(&f.fields, "fields", "1", "extracted columns")
	flag.StringVar(&f.delimiter, "delimiter", "\t", "delimiter")
	flag.BoolVar(&f.separated, "separated", false, "select only with delimiter")

	flag.Parse()

	return f
}

func parseColumns(fields string) []int {
	columnsStr := strings.Split(fields, " ")
	columns := make([]int, len(columnsStr))

	for k, v := range columnsStr {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}

		if n < 1 {
			log.Fatalln("bad fields")
		}

		columns[k] = n - 1
	}

	return columns
}

func cutLine(line string, columns []int, f Flags) string {
	arr := strings.Split(line, f.delimiter)

	if f.separated && len(arr) == 1 {
		return ""
	}

	var ret []string

	for _, v := range columns {
		if len(arr) > v {
			ret = append(ret, arr[v])
		}
	}

	return strings.Join(ret, f.delimiter)
}

func main() {
	f := parseFlags()
	columns := parseColumns(f.fields)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = cutLine(line, columns, f)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
