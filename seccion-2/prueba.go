package main

import (
	"fmt"
)

var channel chan int = make(chan int)

func main() { // ver en playground
	go meEjecutaranEnOtraGorutina()
	go meEjecutaranEnOtraGorutina2()

	fmt.Println("Soy la gorutina principal")
	dato1 := <-channel // recibo un valor del canal
	fmt.Println(dato1)
	dato2 := <-channel // recibo un valor del canal
	fmt.Println(dato2)

}

func meEjecutaranEnOtraGorutina() {
	fmt.Println("Soy otra gorutina")
	channel <- 10 // wg.add
}

func meEjecutaranEnOtraGorutina2() {
	fmt.Println("Soy otra gorutina")
	channel <- 20 // wg.add
}
