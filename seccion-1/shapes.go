package main

import (
	"fmt"
	"math"
)

type medible interface {
	calcularArea() float64
}

type Rectangulo struct {
	base   float64
	altura float64
}

func (r Rectangulo) calcularArea() float64 {
	return r.base * r.altura
}

type Circulo struct {
	Radio float64
}

type Triangulo struct {
	nose float64
}

func (t Triangulo) calcularArea() float64 {
	return 0
}

func (c Circulo) calcularArea() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Rodar() {
	fmt.Println("wiiii")
}

func tienenMismoArea(a medible, b medible) bool {
	return a.calcularArea() == b.calcularArea()
}

func main() {

	var c Circulo = Circulo{
		Radio: 1235,
	}

	var r Rectangulo = Rectangulo{
		base:   123,
		altura: 123,
	}

	c.Rodar()
	fmt.Println(tienenMismoArea(r, c))
}
