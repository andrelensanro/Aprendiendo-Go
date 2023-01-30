package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main(){
	leoArchivo()
	leoArchivo2()
	escribirArchivo()
	escribirArchivo2()
}
func leoArchivo(){
	archivo, err := ioutil.ReadFile("./Hola Mundo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		fmt.Println(string(archivo))
	}
}

func leoArchivo2(){
	archivo, err := os.Open("./Hola Mundo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}
/**Crear un archivo y si ya existe, sobreescribirlo*/
func escribirArchivo(){
	archivo, err := os.Create("./Hola.txt")
	if err!= nil {
        fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una línea nueva")
	archivo.Close()
}
/**Adicionar lineas a un archivo, en este caso, al archivo creado con la primera funcion*/
func escribirArchivo2(){
	filename := "./Hola.txt"
	if Append(filename, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la 2da linea")
    }

}
func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Hubo un error")
        return false
	}
    _, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
        return false
	}
	return true

}


