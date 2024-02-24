package main

import "fmt"

func main() {
	// Operadores aritméticos
	/* var X int = 10
	var Y int = 5 */
	/* suma := X + Y
	resta := X - Y
	multiplicacion := X * Y
	division := X / Y
	modulo := X % Y

	fmt.Println("Operadores aritméticos:")
	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicación:", multiplicacion)
	fmt.Println("División:", division)
	fmt.Println("Módulo:", modulo) */

	/* // Operadores relacionales
	fmt.Println("\nOperadores relacionales:")
	fmt.Println("X == Y:", X == Y) //falso
	fmt.Println("X != Y:", X != Y) //verdadero
	fmt.Println("X > Y:", X > Y)   // true
	fmt.Println("X < Y:", X < Y)   // false
	fmt.Println("X >= Y:", X >= Y) // true
	fmt.Println("X <= Y:", X <= Y) // falso

	if X == Y || X > Y && X == Y {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}
	*/
	/* // Operadores lógicos
	A := 10 == 10 //true
	B := 5 != 5   //false

	A = false
	fmt.Println("\nOperadores lógicos:")
	fmt.Println("A && B:", A && B) // Conjunción (AND) y &
	fmt.Println("A || B:", A || B) // Disyunción (OR)



	fmt.Println("!A:", !A) // Negación */

	// Operadores de asignación
	var Z int
	Z = 16
	fmt.Println("Z =", Z)

	switch Z > 10 {
	case true:
		fmt.Println("es true")
	case false:
		fmt.Println("es false")
	default:
		fmt.Println("??¡¡¡")
	}

}
