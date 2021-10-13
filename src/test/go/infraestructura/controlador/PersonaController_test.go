package controlador

import (
	"bytes"
	"encoding/json"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/aplicacion"
	model "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	repo "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/repositorio"
	controller "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/infraestructura/controlador"
	mongo "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/infraestructura/persistencia/mongo/repositorio"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	persona model.Persona = model.Persona{
		Id:       primitive.NewObjectID(),
		Nombre:   "Jonathan",
		Apellido: "Martinez",
		DireccionCasa: model.Direccion{
			NombreCalle:  "Rivadavia",
			Altura:       1000,
			Piso:         2,
			Departamento: 1,
		},
	}

	personaRepo       repo.PersonaRepositorio      = mongo.NewPersonaRepoMongo()
	personaService    aplicacion.AdmPersonaService = aplicacion.NewAdmPersonaService(personaRepo)
	personaController controller.PersonaController = controller.NewPersonaController(personaService)
)

func TestCrearPersona(t *testing.T) {
	jsonStr, _ := json.Marshal(persona)
	req, _ := http.NewRequest("POST", "/persona", bytes.NewBuffer(jsonStr))
	handler := http.HandlerFunc(personaController.CrearPersona)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.Equal(t, status, http.StatusOK)

	// Cleanup database
	personaRepo.EliminarPersona(persona.Id.String())
}
