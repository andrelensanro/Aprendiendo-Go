package main

import "fmt"
/*Problema del arbolito*/
func main(){
	var n int

	fmt.Scanln(&n)

	for k:=1; k<=n*2; k++{
		for i:=0; i<n*2-k; i++{
			fmt.Printf(".")
		}
		fmt.Printf("*")
		for j:=2; j<=k; j++{
			fmt.Printf(" ")
			fmt.Printf("*")
		}
		for i:=0; i<n*2-k; i++{
			fmt.Printf(".")
		}
		fmt.Println("")
	}
	for i:=0; i<n; i++ {
		for j:=0; j<n*2-1; j++{
			fmt.Printf(".")
		}
		fmt.Printf("*")
		for j:=0; j<n*2-1; j++{
			fmt.Printf(".")
		}
		fmt.Println("")
	}
}