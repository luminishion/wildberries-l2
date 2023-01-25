package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	k int
	n bool
	r bool
	u bool
}

func parseFlags() Flags {
	var f Flags

	flag.IntVar(&f.k, "k", 0, "")
	flag.BoolVar(&f.n, "n", false, "")
	flag.BoolVar(&f.r, "r", false, "")
	flag.BoolVar(&f.u, "u", false, "")

	flag.Parse()

	return f
}

func fileLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var ret []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return ret
}

func parseArgs() string {
	if len(os.Args) < 2 {
		log.Fatalln("wrong arguments count")
	}

	fname := os.Args[len(os.Args)-1]

	return fname
}

func linesCols(lines []string) [][]string {
	ret := make([][]string, len(lines))

	for k, v := range lines {
		ret[k] = strings.Fields(v)
	}

	return ret
}

func main() {
	f := parseFlags()
	fname := parseArgs()
	lines := fileLines(fname)

	if f.u {
		inRet := make(map[string]struct{})
		var ret []string

		for _, v := range lines {
			if _, ok := inRet[v]; ok {
				continue
			}

			inRet[v] = struct{}{}
			ret = append(ret, v)
		}

		lines = ret
	}

	col := linesCols(lines)

	k := f.k

	if !f.n {
		sort.Slice(lines, func(i, j int) bool {
			a := col[i]
			b := col[j]

			if k >= len(a) || k >= len(b) {
				return false
			}

			c := a[k] < b[k]

			if f.r {
				c = !c
			}

			return c
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			a := col[i]
			b := col[j]

			if k >= len(a) || k >= len(b) {
				return false
			}

			na, err := strconv.Atoi(a[k])
			if err != nil {
				return false
			}

			nb, err := strconv.Atoi(b[k])
			if err != nil {
				return false
			}

			c := na < nb

			if f.r {
				c = !c
			}

			return c
		})
	}

	fmt.Println(strings.Join(lines, "\n"))
}
