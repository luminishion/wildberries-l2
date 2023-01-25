package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import (
	"fmt"
)

type FiniteStateMachine struct {
	state State
}

func NewFiniteStateMachine() *FiniteStateMachine {
	s := &StateA{}

	m := &FiniteStateMachine{}
	m.SetState(s)
	return m
}

func (m *FiniteStateMachine) SetState(s State) {
	m.state = s
}

func (m *FiniteStateMachine) Print() {
	m.state.Print()
}

type State interface {
	Print()
}

type StateA struct {
}

func (s *StateA) Print() {
	fmt.Println("A")
}

type StateB struct {
}

func (s *StateB) Print() {
	fmt.Println("B")
}

func main() {
	m := NewFiniteStateMachine()
	m.Print()
	m.SetState(&StateB{})
	m.Print()
}
