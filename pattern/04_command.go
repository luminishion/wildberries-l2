package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import (
	"fmt"
)

type Cmd interface {
	Execute()
}

type TurnOnCmd struct {
	receiver *Receiver
}

func (c *TurnOnCmd) Execute() {
	c.receiver.TurnOn()
}

type TurnOffCmd struct {
	receiver *Receiver
}

func (c *TurnOffCmd) Execute() {
	c.receiver.TurnOff()
}

type Receiver struct {
}

func (s *Receiver) TurnOn() {
	fmt.Println("turn on")
}

func (s *Receiver) TurnOff() {
	fmt.Println("turn off")
}

type Invoker struct {
	queue []Cmd
}

func (i *Invoker) PushCmd(c Cmd) {
	i.queue = append(i.queue, c)
}

func (i *Invoker) Run() {
	for _, v := range i.queue {
		v.Execute()
	}
}

func main() {
	r := &Receiver{}
	i := &Invoker{}

	onCmd := &TurnOnCmd{receiver: r}
	i.PushCmd(onCmd)

	offCmd := &TurnOffCmd{receiver: r}
	i.PushCmd(offCmd)

	i.Run()
}
