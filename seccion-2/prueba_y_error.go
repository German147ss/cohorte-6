package main

//Servidor/servicio/api escrito en Go -> Vamos a consultar una DB externa, de Rick and morty
import (
	"fmt"
	"time"
)

func main() {
	go WithRecover(tareaSecundaria)

	result, err := LlamadaAServicioExterno()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
	time.Sleep(1 * time.Second)
}

func WithRecover(gorutina func()) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	gorutina()
}

func tareaSecundaria() {
	names := []string{
		"lobster",      //0
		"sea urchin",   //1
		"sea cucumber", //2
	}
	fmt.Println("My favorite sea creature is:", names[4])
}

func LlamadaAServicioExterno() (string, error) {
	fmt.Println("Simulamos una llamada a un servicio externo....")

	return "Rick Sanchez", nil
	//return "", objeto
}
