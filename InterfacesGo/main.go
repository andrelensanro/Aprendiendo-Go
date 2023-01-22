package main

import (
	"fmt"
)
//intefaces
type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface{
	ClasificacionVegetal() string
}
// objetos
type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
}

type mujer struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
}
// El hombre por tan solo tener declaradas esas funciones que se declararon en la
// interface Humano.
func (h* hombre) respirar(){
	h.respirando = true
}
func (h* hombre) comer(){
	h.respirando = true
}
func (h* hombre) pensar(){
	h.respirando = true
}
func (h* hombre) sexo() string {
	return "Hombre"
}

func (m *mujer) respirar(){
	m.respirando = true
}
func (m *mujer) comer(){
	m.respirando = true
}
func (m *mujer) pensar(){
	m.respirando = true
}
func (m *mujer) sexo() string {
	return "Mujer"
}


func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un/a %s y ya estoy respirando \n", hu.sexo());
}

func main() {
	Pedro := new(hombre)
	HumanosRespirando(Pedro)


	Maria := new(mujer)
	HumanosRespirando(Maria)
}