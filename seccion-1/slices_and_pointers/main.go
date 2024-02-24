package main

import (
	"fmt"
	bucle "myapp/operadores"
)

var numero2 = 2

// Función que intenta modificar el slice
func modificarSlice(slice []int) {
	slice[0] = 99
	slice = append(slice, 100)
}

func modificarStringPuntero(s *string) {
	*s = "Chau desde puntero"
}

func modificarString(s string) {
	s = "Chau"
}

func main() {

	fmt.Println(bucle.MiNombre)

	// Crear un string
	var s string = "Hola"

	// Llamar a la función con el string
	modificarString(s)

	// Imprimir el string original después de llamar a la función
	fmt.Println("String Original luego de func:", s)

	// Llamar a la función con el string
	modificarStringPuntero(&s)

	fmt.Println("String Original luego de pointer func:", s)
	/*
		// Crear un slice
		originalSlice := []int{1, 2, 3, 4, 5}

		// Llamar a la función con el slice
		modificarSlice(originalSlice)

		// Imprimir el slice original después de llamar a la función
		fmt.Println("Slice Original:", originalSlice)
	*/
}
