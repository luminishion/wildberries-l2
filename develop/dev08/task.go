package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type Command struct {
	cmd    string
	args   []string
	stdin  io.Reader
	stdout io.Writer
}

type Shell struct {
	stdin    io.Reader
	stdout   io.Writer
	handlers map[string]func(*Shell, Command)
}

func NewShell(stdin io.Reader, stdout io.Writer) *Shell {
	s := &Shell{
		stdin:  stdin,
		stdout: stdout,
		handlers: map[string]func(*Shell, Command){
			"cd":   (*Shell).cd,
			"pwd":  (*Shell).pwd,
			"echo": (*Shell).echo,
			"kill": (*Shell).kill,
			"exec": (*Shell).exec,
			"ps":   (*Shell).ps,
			"nc":   (*Shell).nc,
		},
	}

	return s
}

func (s *Shell) cd(c Command) {
	if len(c.args) == 0 {
		return
	}

	err := os.Chdir(c.args[0])
	if err == nil {
		return
	}

	fmt.Fprintln(c.stdout, err)
}

func (s *Shell) pwd(c Command) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(c.stdout, err)
		return
	}

	fmt.Fprintln(c.stdout, path)
}

func (s *Shell) echo(c Command) {
	fmt.Fprintln(c.stdout, strings.Join(c.args, " "))
}

func (s *Shell) kill(c Command) {
	if len(c.args) == 0 {
		return
	}

	pid, err := strconv.Atoi(c.args[0])
	if err != nil {
		fmt.Fprintln(c.stdout, "bad id")
		return
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(c.stdout, "bad id")
		return
	}

	err = proc.Kill()
	if err != nil {
		fmt.Fprintln(c.stdout, "cant kill")
		return
	}
}

func (s *Shell) exec(c Command) {
	if len(c.args) == 0 {
		fmt.Fprintln(c.stdout, "no cmd")
	}

	kostil := append([]string{"/C"}, c.args...)
	cmd := exec.Command("cmd", kostil...)
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(c.stdout, err)
	}
}

func (s *Shell) ps(c Command) {
	l, _ := ps.Processes()
	fmt.Fprintf(c.stdout, "%s\t%s\t%s\n", "pid", "ppid", "nam")

	for _, p := range l {
		fmt.Fprintf(c.stdout, "%d\t%d\t%s\n", p.Pid(), p.PPid(), p.Executable())
	}
}

func (s *Shell) nc(c Command) {
	if len(c.args) < 2 {
		fmt.Fprintln(c.stdout, "usage example: nc tcp google.com:80")
		return
	}

	conn, err := net.Dial(c.args[0], c.args[1])
	if err != nil {
		fmt.Fprintln(c.stdout, "dial failed")
		return
	}
	defer conn.Close()

	go func() {
		_, err := io.Copy(c.stdout, conn)
		if err != nil {
			conn.Close()
		}
	}()
	io.Copy(conn, c.stdin)

	fmt.Fprintln(c.stdout, "closed")
}

func (s *Shell) runCmd(cmd Command) {
	if fn, ok := s.handlers[cmd.cmd]; ok {
		fn(s, cmd)
		return
	}

	fmt.Fprintln(s.stdout, "command not found")
}

func (s *Shell) runBatch(batch []Command) {
	for _, cmd := range batch {
		s.runCmd(cmd)
	}
}

func (s *Shell) connectBatch(batch []Command) {
	if len(batch) == 0 {
		return
	}

	batch[0].stdin = s.stdin

	for i := 1; i < len(batch); i++ {
		var bf bytes.Buffer

		batch[i-1].stdout = &bf
		batch[i].stdin = &bf
	}

	batch[len(batch)-1].stdout = s.stdout

	s.runBatch(batch)
}

func (s *Shell) makeBatch(str string) {
	cmds := strings.FieldsFunc(str, func(i rune) bool {
		return i == '|'
	})

	batch := make([]Command, len(cmds))

	for i, cmd := range cmds {
		arr := strings.Fields(cmd)

		var cmd Command
		cmd.cmd = arr[0]

		if len(arr) > 1 {
			cmd.args = arr[1:]
		}

		batch[i] = cmd
	}

	s.connectBatch(batch)
}

func (s *Shell) Run() error {
	scanner := bufio.NewScanner(s.stdin)

	for scanner.Scan() {
		line := scanner.Text()
		s.makeBatch(line)
	}

	return scanner.Err()
}

func main() {
	s := NewShell(os.Stdin, os.Stdout)
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
