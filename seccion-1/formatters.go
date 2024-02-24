// Golang Program to illustrate the usage of
// range keyword over string in Golang

package main

import "fmt"

// Constructing main function
func main() {

	// tomando un string
	var string = "GeeksforGeeks"

	// usando la palabra clave range con bucle for para
	// iterar sobre la cadena
	for i, item := range string {

		// Imprime el índice de todos los
		// caracteres en la cadena
		fmt.Printf("string[%d] = %d \n", i, item)
	}
}
