package main

import (
	"fmt"
	"strings"
	"time"
)
func main(){

	go miNombreLento("Andrea Selene")
	fmt.Println("Soy la siguiente instruccion despues de llamar a la funcion asincrona, veremos si puede terminar")
	var algo string
	fmt.Scanln(&algo)

}


func miNombreLento(nombre string){
	letras := strings.Split(nombre, "")
	for _, letra := range letras{

		time.Sleep(time.Second)
		fmt.Println(letra)

	}
}