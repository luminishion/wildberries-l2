package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import (
	"fmt"
	"strings"
)

type Handler interface {
	Process(string) string
}

type Lower struct {
	Next Handler
}

func (h *Lower) Process(str string) string {
	str = strings.ToLower(str)

	if h.Next != nil {
		return h.Next.Process(str)
	}

	return str
}

func (h *Lower) SetNext(handler Handler) {
	h.Next = handler
}

type Trim struct {
	Next   Handler
	Cutset string
}

func (h *Trim) Process(str string) string {
	str = strings.Trim(str, h.Cutset)

	if h.Next != nil {
		return h.Next.Process(str)
	}

	return str
}

type Replace struct {
	Next     Handler
	What, To string
}

func (h *Replace) Process(str string) string {
	str = strings.ReplaceAll(str, h.What, h.To)

	if h.Next != nil {
		return h.Next.Process(str)
	}

	return str
}

func main() {
	format := &Trim{
		Next: &Lower{
			Next: &Replace{
				What: "g",
				To:   "o",
			},
		},
		Cutset: " ",
	}

	fmt.Println(format.Process(" 1G2 "))
}
