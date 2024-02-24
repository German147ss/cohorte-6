package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la finalización de esta goroutine al finalizar
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simula un trabajo que toma 1 segundo
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementa el contador de goroutines en espera
		go worker(i, &wg)
	}

	wg.Wait() // Espera hasta que todas las goroutines finalicen

	fmt.Println("All workers have finished")
}
