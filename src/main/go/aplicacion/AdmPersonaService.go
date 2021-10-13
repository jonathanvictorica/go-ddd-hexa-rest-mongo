package aplicacion

import (
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/repositorio"
)

type service struct{}

type AdmPersonaService interface {
	CrearPersona(persona model2.Persona) error
	ModificarPersona(persona model2.Persona) error
	EliminarPersona(id string) error
	BuscarPersonaPorId(id string) (model2.Persona, error)
	BuscarPersonaPorNombre(nombre string) (model2.Persona, error)
	BuscarTodasPersonas() (model2.Personas, error)
}

var (
	personaRepo repositorio.PersonaRepositorio
)

func NewAdmPersonaService(repo repositorio.PersonaRepositorio) AdmPersonaService {
	personaRepo = repo
	return &service{}
}

func (*service) CrearPersona(persona model2.Persona) error {
	return personaRepo.CrearPersona(persona)
}

func (*service) ModificarPersona(persona model2.Persona) error {
	return personaRepo.ModificarPersona(persona)
}

func (*service) EliminarPersona(id string) error {
	return personaRepo.EliminarPersona(id)
}
func (*service) BuscarPersonaPorId(id string) (model2.Persona, error) {
	return personaRepo.BuscarPersonaPorId(id)
}

func (*service) BuscarPersonaPorNombre(nombre string) (model2.Persona, error) {
	return personaRepo.BuscarPersonaPorNombre(nombre)
}

func (*service) BuscarTodasPersonas() (model2.Personas, error) {
	return personaRepo.BuscarTodasPersonas()
}
