package main

import "fmt"

type Ave interface {
	volar()
}

type Gallina struct {
	Peso                   int
	CantidadDeHuevosPorDia int
}

func (g Gallina) Cantar() {
	fmt.Println("KIKIKIIKI")
}

func (g Gallina) volar() {
	fmt.Println("Vuelo KIKIKIKI")
}

func main() {
	var marta Gallina = Gallina{
		Peso:                   20,
		CantidadDeHuevosPorDia: 10,
	}
	fmt.Println(marta)
}
