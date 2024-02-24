package bucle

import (
	"fmt"
)

var MiNombre = "German"

func main() {

	/* //bucle for normal
	for {
		fmt.Println("te voy a saludar....")
		fmt.Println("Hola")
		time.Sleep(3 * time.Second)
	} */

	/* //bucle condition
	i := 0
	for i < 5 {
		fmt.Println(i) // 0  // 1
		i++            // i = 0 -> i = i + 1  -> i = 1 // i = 2
	}

	fmt.Println("valor de i al terminar el bucle")

	fmt.Println(i) */

	/* //bucle for clasico
	for j := 0; j < 5; j++ {
		fmt.Println(j)
		valorAcumuladoDeJ = j
	}

	fmt.Println(valorAcumuladoDeJ) */

	//bucle for range
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	/*
		sharks
		0 = "hammerhead"
		1 = "great white"
		2 = "dogfish"
		.....
	*/

	/* for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	} */

	/*
		for index, value := range sequence {
		      // ...
		}
	*/

	for _, value := range sharks {
		fmt.Println(value)
	}
}
