package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func copyTo(gracefulShutdown chan os.Signal, dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
		gracefulShutdown <- os.Interrupt
	}
}

func buildAddress(args []string) string {
	var b strings.Builder

	b.WriteString(args[0])
	b.WriteString(":")
	b.WriteString(args[1])

	return b.String()
}

func main() {
	fTimeout := flag.Int("t", 10, "Connection end time in seconds")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("error: empty ip and port")
	}

	// Создаем TCP-соединение
	addr := buildAddress(args)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Выход по Ctrl + С
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

	go func() {
		<-time.After(time.Duration(*fTimeout) * time.Second)
		gracefulShutdown <- os.Interrupt
	}()

	// Общение клиента с сервером
	go copyTo(gracefulShutdown, os.Stdout, conn) // читаем из сокета
	go copyTo(gracefulShutdown, conn, os.Stdin)  // пишем в сокет

	<-gracefulShutdown
	log.Println("connection was closed")
}
