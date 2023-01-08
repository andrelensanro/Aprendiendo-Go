package main

import "fmt"

var tabla [10]int

func main(){
	tabla[1]=1
	tabla[4]=15
	fmt.Println(tabla)


	mitabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(mitabla)

	for i:=0; i<len(mitabla); i++{
		fmt.Println(mitabla[i])
		
	}
	for i:=0; i<len(mitabla); i++{
		fmt.Printf("%d ", mitabla[i])
	}
	fmt.Println()
	var matriz [3][2]int // 3 filas 2 columnas 
	matriz[0][0] = 1
	fmt.Println(matriz)


	slice := []int {1, 2, 3}

	fmt.Println(slice)

	variante2()
	variante3()
	variante4()
}


func variante2(){
	elementos := [5]int {1, 2, 3, 4, 5};
	
	slicePy := elementos[2:4]
	fmt.Println(slicePy)
}

func variante3(){
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos) )
	lon := len(elementos)
	//capa := cap(elementos)
	elementos[0] = -1
	fmt.Println(elementos)
	for i:=0; i<100; i++{
		if(i<lon){
			elementos[i] = -1
		}else{ 
			elementos = append(elementos, -3)
		}
	}
	fmt.Println(elementos)
}

func variante4(){
	elementos := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	slice := elementos[:4]
	fmt.Println("slice en seguida de la asignación por refencia", slice)
	for i:=0; i<len(slice); i++{
		slice[i] *= 2
	}
	fmt.Println("slice después de la multiplicación por dos",slice)

	fmt.Println("Modificación del vector elementos por slice quedando así", elementos)
}
