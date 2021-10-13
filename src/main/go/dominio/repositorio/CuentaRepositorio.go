package repositorio

import (
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
)

type CuentaRepositorio interface {
	CrearCuenta(cuenta model2.Cuenta) error
	ModificarCuenta(cuenta model2.Cuenta) error
	EliminarCuenta(id string) error
	BuscarCuentaPorId(id int64) (model2.Cuenta, error)
	BuscarTodasCuenta() (model2.Cuentas, error)
}
