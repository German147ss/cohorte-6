// AHORA POSIBLE IDENTIFICADOR : P120
// AHORA POSIBLE IDENTIFICADOR: B15
package main

type Cochera struct {
	id    string
	horas int
	piso  int
}

type GarageDeGerman [60]Cochera

type GarageDePepito [10]Cochera

/* func (g GarageDeGerman) ObtenerHorasDeCochera(numeroDeCochera int) int {
	return g[numeroDeCochera-1]
} */

const precioPorHora int = 1000

func main() {
	var cochera1 Cochera = Cochera{
		id:    "P115",
		horas: 7,
		piso:  1,
	}

	var garage GarageDeGerman = GarageDeGerman{cochera1, cochera1, cochera1}

}
