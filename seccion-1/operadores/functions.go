package main

import "fmt"

func suma(a int, b int) {
	fmt.Println(MiNombre)
	resultado := a + b
	fmt.Println(resultado)
}

func resta(a int, b int) {
	resultado := a - b
	fmt.Println(resultado)
}

func multiplicacion(a, b int) int {
	return a * b
}

func restaMalaPractica() {
	var X int = 10
	var Y int = 15
	resultado := X - Y
	fmt.Println(resultado)
}

// Se pide sumar / Restar / Multiplicar 2 valores
func main() {
	var X int = 10
	var Y int = 15

	suma(X, Y)
	resta(X, Y)
	fmt.Println(multiplicacion(Y, X))
}
