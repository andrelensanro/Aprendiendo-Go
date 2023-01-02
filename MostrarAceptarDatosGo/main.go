package main

import ( 
	"fmt"
	"os"
	"bufio"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main()  {
	
	fmt.Println("Ingresa el numero 1:")
	fmt.Scanln(&numero1)
	fmt.Println("Ingresa", "el",  "numero", "2:")
	fmt.Scanln(&numero2)
	fmt.Println("Resultado", numero1+numero2)

	fmt.Println("Descripcion:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan(){
		// como no ocurrio ningun error podemos llamar a la función Text
		leyenda = scanner.Text()
	}

	resultado = numero1+numero2

	fmt.Println(leyenda, resultado)
}