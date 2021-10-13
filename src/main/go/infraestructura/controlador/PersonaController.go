package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/aplicacion"
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type personaController struct{}

type PersonaController interface {
	CrearPersona(w http.ResponseWriter, r *http.Request)
	ModificarPersona(w http.ResponseWriter, r *http.Request)
	EliminarPersona(w http.ResponseWriter, r *http.Request)
	BuscarPersonaPorId(w http.ResponseWriter, r *http.Request)
	BuscarPersonaPorNombre(w http.ResponseWriter, r *http.Request)
	BuscarTodasPersonas(w http.ResponseWriter, r *http.Request)
}

var (
	admPersonaService aplicacion.AdmPersonaService
)

func NewPersonaController(service aplicacion.AdmPersonaService) PersonaController {
	admPersonaService = service
	return &personaController{}
}

func (*personaController) CrearPersona(w http.ResponseWriter, r *http.Request) {
	var personaNueva model2.Persona
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}
	json.Unmarshal(reqBody, &personaNueva)
	errPersona := admPersonaService.CrearPersona(personaNueva)
	if errPersona != nil {
		fmt.Fprintf(w, errPersona.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(personaNueva)
}

func (*personaController) ModificarPersona(w http.ResponseWriter, r *http.Request) {
	var personaModificar model2.Persona
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}
	json.Unmarshal(reqBody, &personaModificar)
	errPersona := admPersonaService.ModificarPersona(personaModificar)
	if errPersona != nil {
		fmt.Fprintf(w, errPersona.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(personaModificar)
}

func (*personaController) EliminarPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	admPersonaService.EliminarPersona(vars["id"])
	fmt.Fprintf(w, "Persona eliminada")
}
func (*personaController) BuscarPersonaPorId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	persona, _ := admPersonaService.BuscarPersonaPorId(vars["id"])
	json.NewEncoder(w).Encode(persona)
}

func (*personaController) BuscarPersonaPorNombre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	persona, _ := admPersonaService.BuscarPersonaPorNombre(vars["nombre"])
	json.NewEncoder(w).Encode(persona)
}

func (*personaController) BuscarTodasPersonas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personas, _ := admPersonaService.BuscarTodasPersonas()
	json.NewEncoder(w).Encode(personas)
}
