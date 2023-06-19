package main

import (
	"flag"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

	// Создаем TCP-соединение
	addr := builAddress(args)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()

	// Выход по Ctrl + С
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-time.After(time.Duration(*fTimeout) * time.Second)
		gracefulShutdown <- os.Interrupt
	}()

	// Общение клиента с сервером
	go copyTo(gracefulShutdown, os.Stdout, conn) // читаем из сокета
	go copyTo(gracefulShutdown, conn, os.Stdin)  // пишем в сокет

	<-gracefulShutdown
	logrus.Info("connection was closed")
}

func copyTo(gracefulShutdown chan os.Signal, dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logrus.Info(err)
		gracefulShutdown <- os.Interrupt
	}
}

func builAddress(args []string) string {
	var b strings.Builder

	b.WriteString(args[0])
	b.WriteString(":")
	b.WriteString(args[1])

	return b.String()
}
