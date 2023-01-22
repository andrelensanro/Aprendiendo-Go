package main

import (
	"fmt"
)
type SerVivo interface{
	EstaVivo() bool
}
type humano interface{
	respirar()
	comer()
	pensar()
	sexo() string
	EstaVivo() bool// van a pertenecer a la misma interface de SerVivo
}
type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
	EstaVivo() bool // va a implementar SerVivo también
}

type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p *perro) respirar(){
	p.respirando = true 
}
func (p *perro) comer(){
	p.comiendo = true
}
func (p *perro) EsCarnivoro() bool {
	return p.carnivoro
}
func (p *perro) EstaVivo() bool {
	return p.vivo;
}


type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	vivo bool
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
func (h *hombre) EstaVivo() bool {
	return h.vivo
}

func estoyVivo(sv SerVivo) int {
	if sv.EstaVivo() {return 1}
	return 0
}
func main(){
	Dogo := new(perro)
	Dogo.vivo = true
	Pedro := new(hombre)
	Pedro.vivo = true

	vivos := [2]SerVivo {Dogo, Pedro}
	cont := 0
	for i:=0; i<len(vivos); i++{
		if(vivos[i].EstaVivo()){ cont++;}
	}
	fmt.Println("Hay",cont,"seres vivos")

}