package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import (
	"fmt"
	"math"
)

type Visitor interface {
	VisitTriangle(*Triangle)
	VisitCircle(*Circle)
}

type AreaCalculator struct {
	Sum float64
}

func (v *AreaCalculator) VisitTriangle(t *Triangle) {
	a, b, c := t.A, t.B, t.C
	p := (a + b + c) / 2
	v.Sum += math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func (v *AreaCalculator) VisitCircle(c *Circle) {
	v.Sum += math.Pi * math.Pow(c.Rad, 2)
}

type Shape interface {
	Accept(Visitor)
}

type Triangle struct {
	A, B, C float64
}

func (t *Triangle) Accept(v Visitor) {
	v.VisitTriangle(t)
}

type Circle struct {
	Rad float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

func main() {
	shps := []Shape{
		&Triangle{1, 1, 1},
		&Circle{3},
	}

	areaCalculator := &AreaCalculator{}
	for _, v := range shps {
		v.Accept(areaCalculator)
	}

	fmt.Println(areaCalculator.Sum)
}
