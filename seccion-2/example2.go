package main

import (
	"fmt"
)

var canal chan bool = make(chan bool)

func main() { // hilo principal
	go meEjecutaranEnOtraGorutina()
	fmt.Println("Soy la gorutina principal")
	valorQueLlegoDesdeElCanal := <-canal // recibo un valor del canal / quedo a la esperar de que alguien me mande un valor
	fmt.Println("FIN main", valorQueLlegoDesdeElCanal)
}

func meEjecutaranEnOtraGorutina() {
	fmt.Println("Soy otra gorutina")
	canal <- true // envío un valor al canal
	for i := 1; i <= 10000000000; i++ {

	}
	fmt.Println("FIN meEjecutaranEnOtraGorutina")
}
