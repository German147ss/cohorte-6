package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // Espera un segundo entre cada número
	}
}

func main() {
	go printNumbers()           // Inicia una goroutine para imprimir números
	time.Sleep(5 * time.Second) // Espera 5 segundos para dar tiempo a la goroutine
	fmt.Println("Done")
}
