package main

import "fmt"

func main() {
	exponenteDe(2)
}

func exponenteDe(numero int){
	if numero > 1000000 {
		return
	}
	fmt.Println(numero)
	exponenteDe(numero*2)
}