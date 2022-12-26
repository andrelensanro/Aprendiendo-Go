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


/*
Cuando se crea una nueva variable las inicializa en 0, false
Un := crea y asigna
Este lenguaje considera mucho el poder y posibilidades de la sintaxis:
	Una variable si comienza en mayusculas puede ser accedido desde otro paquete
	Una variable si comienza en minuscula es local
	Lo mismo en con las funciones
No se puede convertir un tipo de dato a otro tipo de dato, int32 a int64 tampoco es posible
Los tipos de datos tambien son funciones, se usa para convetir

*/
package main
import (
	"fmt"
)


var status bool = false 


func main(){
	numero2, numero3, numero4, texto, status:= 2, 3, 67, "este es mi segundo programa", true
	fmt.Println("Linea 1: ", numero3)
	fmt.Println("Linea 2: ", numero2)
	fmt.Println("Linea 3: ", numero4)
	fmt.Println("Linea 4: ", texto)
	fmt.Println("Linea 5: ", status)
	var moneda float32 = 126.67
	fmt.Println("Linea 6: ", moneda)
	numero2 = int(moneda)
	// aun no estoy muy segura si esa conversion de int(moneda de la linea de arriba hace la conversion de moneda)
	texto = fmt.Sprintf("Linea 7: %f", moneda) // use %.2f
	fmt.Println(texto)
	fmt.Println("Linea 8: ", moneda)
	fmt.Println("Linea 9: ", numero2)
	cosaScope()
}

func cosaScope(){
	fmt.Println("Linea 10: ", status)
}
