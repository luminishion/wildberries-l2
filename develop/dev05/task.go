package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func parseFlags() Flags {
	var f Flags

	flag.IntVar(&f.after, "after", 0, "печатать +N строк после совпадения")
	flag.IntVar(&f.before, "before", 0, "печатать +N строк до совпадения")
	flag.IntVar(&f.context, "context", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.count, "count", false, "количество строк")
	flag.BoolVar(&f.ignoreCase, "ignoreCase", false, "игнорировать регистр")
	flag.BoolVar(&f.invert, "invert", false, "вместо совпадения, исключать")
	flag.BoolVar(&f.fixed, "fixed", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&f.lineNum, "lineNum", false, "печатать номер строки")

	flag.Parse()

	return f
}

func parseArgs() (string, string) {
	if len(os.Args) < 3 {
		log.Fatalln("wrong arguments count")
	}

	fname := os.Args[len(os.Args)-1]
	pattern := os.Args[len(os.Args)-2]

	return fname, pattern
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

func findR(lines []string, pattern string) []int {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalln("bad pattern")
	}

	var ret []int

	for k, v := range lines {
		if reg.MatchString(v) {
			ret = append(ret, k)
		}
	}

	return ret
}

func findF(lines []string, pattern string) []int {
	var ret []int

	for k, v := range lines {
		if strings.Index(v, pattern) != -1 {
			ret = append(ret, k)
		}
	}

	return ret
}

func selectLines(f Flags, length int, idx []int) []int {
	a := f.context
	b := f.context

	if f.before != 0 {
		b = f.before
	}

	if f.after != 0 {
		a = f.after
	}

	var ret []int

	for _, v := range idx {
		start := v - a
		if start < 0 {
			start = 0
		}

		end := v + b + 1
		if end > length {
			end = length
		}

		for i := start; i < end; i++ {
			ret = append(ret, i)
		}
	}

	return ret
}

func invert(length int, idx []int) []int {
	var ret []int
	c := 0

	for i := 0; i < length; i++ {
		if len(idx) == 0 || c >= len(idx) || idx[c] != i {
			ret = append(ret, i)
		} else {
			c++
		}
	}

	return ret
}

func main() {
	f := parseFlags()
	fname, pattern := parseArgs()
	input := fileLines(fname)

	var lines []string
	if f.ignoreCase {
		lines = make([]string, len(input))
		for k, v := range input {
			lines[k] = strings.ToLower(v)
		}
	} else {
		lines = input
	}

	var idx []int
	if f.fixed {
		idx = findF(lines, pattern)
	} else {
		idx = findR(lines, pattern)
	}

	if f.count {
		fmt.Println(len(idx))
		return
	}

	if f.invert {
		idx = invert(len(lines), idx)
	} else {
		idx = selectLines(f, len(lines), idx)
	}

	if f.lineNum {
		for _, v := range idx {
			fmt.Println(v+1, input[v])
		}
	} else {
		for _, v := range idx {
			fmt.Println(input[v])
		}
	}
}
