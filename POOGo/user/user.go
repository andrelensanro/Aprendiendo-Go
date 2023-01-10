package user

import "time"

type Usuario struct{
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
}

// this hace referencia a la estructura en la que estoy parado

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool){
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}