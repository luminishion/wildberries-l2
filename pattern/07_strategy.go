package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import (
	"fmt"
	"math/rand"
	"time"
)

type StrategyRand interface {
	Get() uint32
}

type LcgRand struct {
	val uint32
}

func NewLcgRand() *LcgRand {
	return &LcgRand{
		val: uint32(time.Now().UnixNano()),
	}
}

func (r *LcgRand) Get() uint32 {
	r.val = (1103515245*r.val + 12345) % 2147483648
	return r.val
}

type GoRand struct {
	rnd *rand.Rand
}

func NewGoRand() *GoRand {
	return &GoRand{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *GoRand) Get() uint32 {
	return r.rnd.Uint32()
}

type Context struct {
	strategy StrategyRand
}

func (c *Context) Algorithm(a StrategyRand) {
	c.strategy = a
}

func (c *Context) Get() uint32 {
	return c.strategy.Get()
}

func main() {
	ctx := &Context{}

	lcg := NewLcgRand()
	ctx.Algorithm(lcg)
	fmt.Println(ctx.Get())

	gorand := NewGoRand()
	ctx.Algorithm(gorand)
	fmt.Println(ctx.Get())
}
