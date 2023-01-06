package main

import "fmt"

func main(){
	var entrada int
	fmt.Println("Ingresa un valor entero:")
	fmt.Scanln(&entrada)
	fmt.Println("Numero", entrada, "por dos es", uno(entrada))
	numero, estado := dos(entrada)
	fmt.Println("Funcion con dos valores retornados:", numero, estado)
	z := nombrandoVarRetorno(entrada)
	fmt.Println("Funcion que nombra la variable de retorno. El valor de z es", z)

	
	fmt.Println(Calculo(2))
	fmt.Println(Calculo(20,3))
	fmt.Println(Calculo(2,31,5,6))
	fmt.Println(Calculo(12, 34, 21))
}

func uno(numero int) int {
	return numero*2
}
func dos(numero int) (int, bool) {
	if numero == 1{
		return numero, false
	}
	return numero, true 
}
func nombrandoVarRetorno(numero int) (z int){
	z = numero * 3 
	return
}

func Calculo(numero ...int) int{
	total:=0
	//for it, num := range numero {
	for _, num := range numero{
		total = total + num
		//fmt.Println(it, num)
	} 
	return total
}
