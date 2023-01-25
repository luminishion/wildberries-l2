package main

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

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func parseFlags() *string {
	timeoutStr := flag.String("timeout", "10s", "connection timeout")

	flag.Parse()

	return timeoutStr
}

func parseArgs() (string, string) {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		log.Fatalln("wrong arguments count")
	}

	port := os.Args[len(os.Args)-1]
	ip := os.Args[len(os.Args)-2]

	return ip, port
}

type Telnet struct {
	con net.Conn
}

func NewTelnet() *Telnet {
	return &Telnet{}
}

func (t *Telnet) Connect(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}

	t.con = conn
	return nil
}

func (t *Telnet) Read() error {
	_, err := io.Copy(t.con, os.Stdin)
	return err
}

func (t *Telnet) Write() error {
	_, err := io.Copy(os.Stdout, t.con)
	return err
}

func (t *Telnet) Close() error {
	return t.con.Close()
}

func main() {
	ip, port := parseArgs()
	addr := net.JoinHostPort(ip, port)

	timeoutStr := parseFlags()
	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		log.Fatalln("time parse error")
	}

	cl := NewTelnet()

	if err := cl.Connect(addr, timeout); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := cl.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	closed := make(chan error, 2)

	go func() {
		closed <- cl.Read()
	}()

	go func() {
		closed <- cl.Write()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-quit:
	case err := <-closed:
		if err != nil {
			log.Fatalln(err)
		}
	}
}
