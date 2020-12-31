package models

type Persona struct {
	Nombre          string           `json:"nombre"`
	Apellido        string           `json:"apellido"`
	Edad            int              `json:"edad"`
	DatosPersonales *DatosPersonales `json:"datos_personales"`
}

func (p *Persona) MayorDeEdad() bool {
	return p.Edad >= 18
}

func (p *Persona) CambiarNombre(nombre string) {
	p.Nombre = nombre
}

type DatosPersonales struct {
	Direccion string `json:"direccion"`
	Telefono  int    `json:"telefono"`
}
