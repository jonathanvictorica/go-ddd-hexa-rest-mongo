package model

import "math/big"

type Cuenta struct {
	Id         int64
	TipoCuenta string
	Saldo      big.Int
}

type Cuentas []Cuenta
