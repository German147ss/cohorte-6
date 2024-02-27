package main

import (
	"fmt"
	"sync"
	"time"
)

func sequentialSum() {
	sum := 0
	start := time.Now() //
	for i := 1; i <= 10000000; i++ {
		sum += i * i
	}
	elapsed := time.Since(start) //
	fmt.Printf("Secuencial: La suma de los cuadrados es: %d, Tiempo de ejecución: %s\n", sum, elapsed)
}

func concurrentSum() {
	var wg sync.WaitGroup // Creamos una WaitGroup para esperar que todas las goroutines finalicen
	sum := 0              // Variable para almacenar la suma de los cuadrados
	start := time.Now()   // Registramos el tiempo de inicio

	var ch chan int = make(chan int) // Creamos un canal para comunicarnos con las goroutines

	for i := 1; i <= 10000000; i++ {
		wg.Add(1)        // Añadimos 1 al contador de goroutines esperadas en la WaitGroup
		go func(x int) { // Iniciamos una goroutine para calcular el cuadrado de un número
			defer wg.Done() // Avisamos a la WaitGroup que hemos terminado al salir de la goroutine
			ch <- x * x     // Enviamos el resultado al canal
			return
		}(i) // Pasamos el valor de i a la goroutine
	}

	go func() {
		wg.Wait() // Esperamos a que todas las goroutines hayan terminado
		close(ch) // Cerramos el canal para indicar que ya no se enviarán más valores
	}()

	//
	for result := range ch { // Iteramos sobre los valores recibidos del canal hasta que esté cerrado // for ......1 ....... 4
		sum += result // Sumamos cada valor recibido al total
	}

	elapsed := time.Since(start)                                                                        // Calculamos el tiempo transcurrido
	fmt.Printf("Concurrente: La suma de los cuadrados es: %d, Tiempo de ejecución: %s\n", sum, elapsed) // Imprimimos el resultado y el tiempo transcurrido
}

func main() {
	fmt.Println("Inicio primera")
	sequentialSum()
	fmt.Println("fin primera")

	time.Sleep(2 * time.Second)
	fmt.Println("Inicio 2da")

	concurrentSum()
	fmt.Println("Fin 2da")

}
