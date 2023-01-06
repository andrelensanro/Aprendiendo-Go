package main

import "fmt"

func main()  {
	i:= 1
	for i <= 10{
		fmt.Println(i)
		i++
	}
	for i:=1; i<=10; i++{
		fmt.Println(i)
	}
	numero:=0
	for{
		fmt.Println("Continuo\nIngresa el número secreto")
		fmt.Scanln(&numero)
		if numero==1234{
			fmt.Println("Has acertado")
			break
		}
	}
	i=0

	RUTINA:
		for i<10{
			if i==4{
				i=i+2
				fmt.Println("Voy a rutina sumando 2 a i")
				goto RUTINA
			}
			fmt.Printf("Valor de i: %d\n", i)
			i++
		}
}
