package main

import (
	"fmt"
	"sort"
)

func main() {
	miMapa := map[string]int{
		"uno":    1,
		"dos":    2,
		"tres":   3,
		"cuatro": 4,
		"cinco":  5,
	}

	fmt.Println(miMapa)

	miMapa2 := map[int]int{
		1: 10,
		2: 200,
		3: 3000,
		4: 40000,
		5: 500000,
	}

	fmt.Println(miMapa2)

	// Crear un slice de estructuras para almacenar pares clave-valor
	var pares []struct {
		Clave string
		Valor int
	}

	// Agregar elementos al slice
	for clave, valor := range miMapa {
		pares = append(pares, struct {
			Clave string
			Valor int
		}{clave, valor})
	}

	// Ordenar el slice por clave
	sort.Slice(pares, func(i, j int) bool {
		return pares[i].Clave < pares[j].Clave
	})

	// Iterar sobre el slice ordenado
	for _, par := range pares {
		fmt.Printf("%s: %d\n", par.Clave, par.Valor)
	}
	var diasSemana map[int]string
	diasSemana[1] = "Lunes"
	fmt.Println(diasSemana)
}
