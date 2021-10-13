package root

import (
	"fmt"
	aplicacion2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/aplicacion"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/repositorio"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/infraestructura/controlador"
	repositorio2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/infraestructura/persistencia/mongo/repositorio"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"

	"log"
	_ "log"
	"net/http"
)

var (
	personaRepoMongo  repositorio.PersonaRepositorio = repositorio2.NewPersonaRepoMongo()
	admPersonaService aplicacion2.AdmPersonaService  = aplicacion2.NewAdmPersonaService(personaRepoMongo)
	personaController controller.PersonaController   = controller.NewPersonaController(admPersonaService)
)

func CreateRoot() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/persona", personaController.CrearPersona).Methods("POST")
	r.HandleFunc("/persona/{id}", personaController.ModificarPersona).Methods("PUT")
	r.HandleFunc("/persona/{id}", personaController.EliminarPersona).Methods("DELETE")
	r.HandleFunc("/persona/{id}", personaController.BuscarPersonaPorId).Methods("GET")
	r.HandleFunc("/persona/nombre/{nombre}", personaController.BuscarPersonaPorNombre).Methods("GET")
	r.HandleFunc("/personas", personaController.BuscarTodasPersonas).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}
