package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/aplicacion"
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type cuentaController struct{}

type CuentaController interface {
	CrearCuenta(w http.ResponseWriter, r *http.Request)
	ModificarCuenta(w http.ResponseWriter, r *http.Request)
	EliminarCuenta(w http.ResponseWriter, r *http.Request)
	BuscarCuentaPorId(w http.ResponseWriter, r *http.Request)
	BuscarTodasCuentas(w http.ResponseWriter, r *http.Request)
}

var (
	admCuentaService aplicacion.AdmCuentaService
)

func NewCuentaController(service aplicacion.AdmCuentaService) CuentaController {
	admCuentaService = service
	return &cuentaController{}
}

func (*cuentaController) CrearCuenta(w http.ResponseWriter, r *http.Request) {
	var CuentaNueva model2.Cuenta
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}
	json.Unmarshal(reqBody, &CuentaNueva)
	errCuenta := admCuentaService.CrearCuenta(CuentaNueva)
	if errCuenta != nil {
		fmt.Fprintf(w, errCuenta.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CuentaNueva)
}

func (*cuentaController) ModificarCuenta(w http.ResponseWriter, r *http.Request) {
	var CuentaModificar model2.Cuenta
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}
	json.Unmarshal(reqBody, &CuentaModificar)
	errCuenta := admCuentaService.ModificarCuenta(CuentaModificar)
	if errCuenta != nil {
		fmt.Fprintf(w, errCuenta.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CuentaModificar)
}

func (*cuentaController) EliminarCuenta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	admCuentaService.EliminarCuenta(vars["id"])
	fmt.Fprintf(w, "Cuenta eliminada")
}
func (*cuentaController) BuscarCuentaPorId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idEntero, _ := strconv.ParseInt(vars["id"], 0, 0)
	Cuenta, _ := admCuentaService.BuscarCuentaPorId(idEntero)
	json.NewEncoder(w).Encode(Cuenta)
}

func (*cuentaController) BuscarTodasCuentas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Cuentas, _ := admCuentaService.BuscarTodasCuentas()
	json.NewEncoder(w).Encode(Cuentas)
}
