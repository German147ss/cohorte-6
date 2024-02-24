package main

import (
	"fmt"
	"time"
)

func enviarMensaje(canal chan string) {
	time.Sleep(2 * time.Second)
	// Enviar un mensaje a través del canal
	canal <- "Hola desde la goroutine"
}

func main() {
	// Crear un canal de tipo string
	canal := make(chan string)

	// Iniciar la goroutine y pasar el canal como argumento
	go enviarMensaje(canal)

	// Leer el mensaje enviado a través del canal
	mensajeRecibido := <-canal
	fmt.Println(mensajeRecibido)
}
