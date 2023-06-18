package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

/*
Реализовать простейший telnet-клиент
Пример вызова: go-telnet --timeout=10s host port go-telnet mysite.ru 8080

Требования:
1. Программа должна подключаться к указанному хосту (ip или
доменное имя + порт) по протоколу TCP. После подключения
STDIN программы должен записываться в сокет, а данные
полученные из сокета должны выводиться в STDOUT
2. Опционально в программу можно передать таймаут на
подключение к серверу (через аргумент --timeout, по
умолчанию 10s)
3. При нажатии Ctrl+D программа должна закрывать сокет и
завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться. При подключении к
несуществующему сервер, программа должна завершаться
через timeout
*/

func main() {
	fTimeout := flag.Int("t", 10, "Connection end time")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		logrus.Error("error: empty ip/domain name or/and port")
		os.Exit(1)
	}

	addr := builAddress(args)
	timeout := time.Duration(*fTimeout) * time.Second
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()

	// Выход по Ctrl + D
	sgChan := make(chan os.Signal, 1)
	signal.Notify(sgChan, os.Interrupt)
	go func() {
		for range sgChan {
			logrus.Info("Received an interrupt, stopping services...")
			conn.Close()
			os.Exit(0)
		}
	}()

	// Общение клиента с сервером
	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)
	for {
		fmt.Print(">> ")
		sendText, _ := clientReader.ReadString('\n')
		_, err := fmt.Fprintln(conn, sendText)
		if err != nil {
			logrus.Info("timeout is over")
			return
		}

		rcvText, err := serverReader.ReadString('\n')
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Print(rcvText)
	}
}

// Принимает слайс строк, где первая строка ip/host,
// вторая строка port и возвращает в виде: "[ip/host]:[port]"
func builAddress(args []string) string {
	var b strings.Builder

	b.WriteString(args[0])
	b.WriteString(":")
	b.WriteString(args[1])

	return b.String()
}
