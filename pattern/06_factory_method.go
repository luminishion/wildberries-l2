package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import (
	"fmt"
)

type ResponseCreator interface {
	CreateResponse(bool) Response
}

type Response interface {
	String() string
	Code() int
}

type ResponseOk struct {
}

func (r *ResponseOk) String() string {
	return "OK"
}

func (r *ResponseOk) Code() int {
	return 200
}

type ResponseBad struct {
}

func (r *ResponseBad) String() string {
	return "err"
}

func (r *ResponseBad) Code() int {
	return 500
}

type ResponseFactory struct {
}

func (f *ResponseFactory) CreateResponse(isOk bool) Response {
	if isOk {
		return &ResponseOk{}
	}

	return &ResponseBad{}
}

func NewResponseCreator() ResponseCreator {
	return &ResponseFactory{}
}

func main() {
	f := NewResponseCreator()

	fmt.Println(f.CreateResponse(true).Code())
	fmt.Println(f.CreateResponse(false).Code())
}
