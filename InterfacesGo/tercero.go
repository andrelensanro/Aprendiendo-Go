package main

import "fmt"

type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
}

type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
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

func AnimalesRespirando(an animal){
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {return 1}
	return 0
}

func main(){
	animalesCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Pancita := new(perro)
	Pancita.carnivoro = false
	Lucas := new(perro)
	Lucas.carnivoro = false

	mis_animales := [3]animal {Dogo, Pancita, Lucas}
	for i:=0; i<len(mis_animales); i++ {
		animalesCarnivoros+=AnimalesCarnivoros(mis_animales[i])
	}
	fmt.Println("Tengo en total", animalesCarnivoros , "animales carnivoros")
}