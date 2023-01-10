package escuela

type Escuela struct {
	Nombre string
	Capacidad int
}
type CentroMedico struct {
	Nombre string
	Especialidad string
}

func (this *Escuela)AltaEscuela(nombre string, capacidad int){
	this.Nombre = nombre
	this.Capacidad = capacidad
}

func (this *CentroMedico) AltaCentroMedico (nombre string, especialidad string){
	this.Nombre = nombre
	this.Especialidad = especialidad
}