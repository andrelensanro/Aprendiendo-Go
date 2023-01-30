/*
package main

import "fmt"

func main() {
    slice := []int{3, 4, 5, 6, 8}
    fmt.Println(slice[3]) // 6
    fmt.Println(slice[7]) // panic: runtime error: index out of range

    fmt.Println("Fin") // no se ejecuta
}

*/

package main

import (
    "fmt"
    "strconv"
)

func main() {
    var cadena = "j"
    numero, err := strconv.Atoi(cadena)
    if err != nil {
        panic(err)
    }
    fmt.Println(numero + 10)
}