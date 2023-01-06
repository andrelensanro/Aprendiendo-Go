package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int{
	return num1 + num2
} 

func main(){
	var num1, num2 int 
	fmt.Scanln(&num1, &num2)
	fmt.Println("La suma de", num1 ,"y", num2, "es", Calculo(num1, num2))
	Calculo = func(num1 int, num2 int) int{
		return num1 - num2
	}
	fmt.Println("La resta de", num1 ,"y", num2, "es", Calculo(num1, num2))
	Calculo = func(num1 int, num2 int) int{
		return num1 * num2
	}
	fmt.Println("La multiplicacion de", num1, "y", num2, "es", Calculo(num1, num2))
	
	
	Operaciones()


	tablaDel2:=2
	Mitabla:=Tabla(tablaDel2) // se convierte en una función que se la pasará ejecutando el 1er return 
	// es por ello que secuencia no se inicializa cada que se manda a llamar la funcion

	/*
	Mitabla se convierte en esto, donde secuencia lo toma del exterior que es una de las caracteristicas de los closures

	func Mitabla() {
		secuencia+=1  
		return numero*secuencia
	}
	*/

	for i:=1; i<=10; i++{
		resultado := Mitabla()
		fmt.Println(resultado)
	} 


}

func Operaciones(){
	resultado := func() int{
		return 23+76
	}
	fmt.Println(resultado())
}


func Tabla(valor int) func() int{
	numero:=valor
	secuencia:=0
	return func() int{
		secuencia+=1  
		return numero*secuencia
	}
}