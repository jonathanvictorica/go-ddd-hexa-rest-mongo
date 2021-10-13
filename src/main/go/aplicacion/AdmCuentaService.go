package aplicacion

import (
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/repositorio"
)

type cuentaService struct{}

type AdmCuentaService interface {
	CrearCuenta(Cuenta model2.Cuenta) error
	ModificarCuenta(Cuenta model2.Cuenta) error
	EliminarCuenta(id string) error
	BuscarCuentaPorId(id int64) (model2.Cuenta, error)
	BuscarTodasCuentas() (model2.Cuentas, error)
}

var (
	CuentaRepo repositorio.CuentaRepositorio
)

func NewAdmCuentaService(repo repositorio.CuentaRepositorio) AdmCuentaService {
	CuentaRepo = repo
	return &cuentaService{}
}

func (*cuentaService) CrearCuenta(Cuenta model2.Cuenta) error {
	return CuentaRepo.CrearCuenta(Cuenta)
}

func (*cuentaService) ModificarCuenta(Cuenta model2.Cuenta) error {
	return CuentaRepo.ModificarCuenta(Cuenta)
}

func (*cuentaService) EliminarCuenta(id string) error {
	return CuentaRepo.EliminarCuenta(id)
}
func (*cuentaService) BuscarCuentaPorId(id int64) (model2.Cuenta, error) {
	return CuentaRepo.BuscarCuentaPorId(id)
}

func (*cuentaService) BuscarTodasCuentas() (model2.Cuentas, error) {
	return CuentaRepo.BuscarTodasCuenta()
}
