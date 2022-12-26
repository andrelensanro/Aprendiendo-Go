/*
Notas: go nos forza a llevar buenas practicas 
A pesar de que es un lenguaje compilado no me devuelve un .exe para ejecutarlo lo hace internamente

Comando para ejecutar seguido de una ejecucion si es que todo esta bien
go run main.go 
Comamndo para ejecutar y sacar el .exe, alta portabilidad para go
go build main.go
*/
/*
package main

import (
	"fmt" // para mostrar  por pantalla
)
// también es valido import main() en una sola linea siempre y cuando sea una sola

func main()  { // declaracion de una funcion
	fmt.Println("Hola Mundo")
}

*/
//-------------------------------------------------------------------------------------------------
package main
import fmt

func main(){

	var numero2 int
	var numero3 int
	var numero4 int
	var texto string 
	numero2 := 2
	numero3 := 3
	numero4 := 67
	texto := "Este es mi texto"
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
}
