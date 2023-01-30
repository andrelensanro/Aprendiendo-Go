package main

import (
	//"fmt"
	//"io/ioutil"
	//"os"
	"log"
)


func main() {
	/*
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("Error opening archivo")
		os.Exit(1)
	}
	*/
	ejemploPanic()
}

func ejemploPanic() {
	defer func(){
		reco := recover()
		if reco!= nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a:=1
	if a==1 {
		// Aborta el programa
		panic("se encontro el valor de 1")
	}
}



