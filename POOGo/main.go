package main

import (
	"fmt"
	"time"
	us "mimodulo/user"
	es "mimodulo/escuela"
)
type usuario struct{
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
}

type pepe struct {
	us.Usuario
}

func main(){
	u := new(pepe)
	u.AltaUsuario(1, "Pablo Tinoco", time.Now(), true)
	fmt.Println(u)
	e := new(es.Escuela)
	e.AltaEscuela("ESCOM", 1029)
	fmt.Println(e)

	cm := new(es.CentroMedico)
	cm.AltaCentroMedico("LA RAZA", "Todas")
	fmt.Println(cm)

}