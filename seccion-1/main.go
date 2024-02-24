package main

import "fmt"

type persona struct {
	Id            int
	Edad          int
	Apellido      string
	Nombre        string
	EsMayorDeEdad bool
}

func (p *persona) setEdad(edad int) {
	p.Edad = edad
}

func (p persona) saludarAMiMismo() {
	fmt.Println("hola ", p.Nombre)
}

func main() {
	var personaGerman persona

	personaGerman.Id = 1234567
	//personaGerman.Edad = 25
	personaGerman.setEdad(25)
	personaGerman.Apellido = "Apellido"
	personaGerman.Nombre = "German"

	var personaLudmila persona

	personaLudmila.Id = 987654
	personaLudmila.Edad = 9
	personaLudmila.Nombre = "Ludmila"

	personaLudmila.saludarAMiMismo()
	fmt.Println(personaGerman)
}
