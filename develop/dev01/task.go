package main

// Программа, печатающая текущее точное время
// используя библиотеку github.com/beevik/ntp
// Программа оформлена как go module
// Программа корректно обрабатывает ошибки билиотеки:
// выводит их в STDERR и возвращает ненулевой код
// выхода в OS

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	h, m, s := time.Clock()
	fmt.Printf("Current time from NTP server: %d:%d:%d\n", h, m, s)
}
