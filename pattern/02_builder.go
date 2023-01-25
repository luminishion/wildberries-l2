package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import (
	"fmt"
	"strings"
)

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.SetName()
	d.builder.SetInfo()
	d.builder.SetPrice()
}

type Product struct {
	data string
}

func (p *Product) GetData() string {
	return p.data
}

type Builder interface {
	SetName()
	SetInfo()
	SetPrice()
	GetResult() *Product
}

type CocaColaBuilder struct {
	sb strings.Builder
}

func (b *CocaColaBuilder) SetName() {
	b.sb.WriteString("<name>CocaCola</name>\n")
}

func (b *CocaColaBuilder) SetInfo() {
	b.sb.WriteString("<info>original one</info>\n")
}

func (b *CocaColaBuilder) SetPrice() {
	b.sb.WriteString("<price>4$</price>\n")
}

func (b *CocaColaBuilder) GetResult() *Product {
	return &Product{b.sb.String()}
}

type DobriyColaBuilder struct {
	sb strings.Builder
}

func (b *DobriyColaBuilder) SetName() {
	b.sb.WriteString("<name>Dobriy Cola</name>\n")
}

func (b *DobriyColaBuilder) SetInfo() {
	b.sb.WriteString("<info>importozameshenie</info>\n")
}

func (b *DobriyColaBuilder) SetPrice() {
	b.sb.WriteString("<price>120RUB</price>\n")
}

func (b *DobriyColaBuilder) GetResult() *Product {
	return &Product{b.sb.String()}
}

func main() {
	colaBuilder := &CocaColaBuilder{}
	colaDirector := &Director{colaBuilder}
	colaDirector.Construct()
	cola := colaBuilder.GetResult()

	dobriyBuilder := &DobriyColaBuilder{}
	dobroDirector := &Director{dobriyBuilder}
	dobroDirector.Construct()
	dobro := dobriyBuilder.GetResult()

	fmt.Println(cola.GetData())
	fmt.Println(dobro.GetData())
}
