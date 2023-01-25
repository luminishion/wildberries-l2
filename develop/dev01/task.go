package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const (
	FORMAT = time.RFC3339
	NTP    = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	tm, err := ntp.Time(NTP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ntp error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("ntp time:", tm.Format(FORMAT))
}
