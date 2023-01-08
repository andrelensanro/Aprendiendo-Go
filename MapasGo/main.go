package main

import "fmt"

func main(){
	paises := make(map[string]string)
	paises["Mexico"] = "CDMX"
	paises["Argentina"] = "Buenos aires"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":39,
		"Real madrid":38,
		"Chivas":37,
		"Boca Juniors":12}

	fmt.Println(campeonato)
	
	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55
	fmt.Println(campeonato)
	delete(campeonato, "Real madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range(campeonato){
		fmt.Println(equipo, puntaje)
	}
	puntaje, existe:= campeonato["Chivas"]
	if(existe){
		fmt.Println("Chivas sí existe", puntaje)
	}else{
		fmt.Println("Chivas no existe", puntaje)
	}
	puntaje, existe = campeonato["Mineria"]
	if(existe){
		fmt.Println("Mineria sí existe", puntaje)
	}else{
		fmt.Println("Mineria no existe", puntaje)
	}
}