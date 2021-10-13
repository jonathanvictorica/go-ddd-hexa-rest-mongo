package repositorio

import (
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
)

type PersonaRepositorio interface {
	CrearPersona(persona model2.Persona) error
	ModificarPersona(persona model2.Persona) error
	EliminarPersona(id string) error
	BuscarPersonaPorId(id string) (model2.Persona, error)
	BuscarPersonaPorNombre(nombre string) (model2.Persona, error)
	BuscarTodasPersonas() (model2.Personas, error)
}
