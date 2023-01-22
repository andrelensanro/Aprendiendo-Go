package main

import "fmt"

type humano interface{
	respirar()
	comer()
	pensar()
	sexo() string
}

type hombre struct{
	respirando bool
	comiendo bool
	pensando bool
	edad int 
	esHombre bool
}

type mujer struct{
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre { return "Hombre" }
	return "Mujer"
}
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Estoy respirando y soy %s \n", hu.sexo())
}


func main(){
	Pedro := new(hombre)
	Pedro.esHombre = true

	HumanosRespirando(Pedro)

	Maria := new(mujer)
	Maria.esHombre = false

	HumanosRespirando(Maria)
}