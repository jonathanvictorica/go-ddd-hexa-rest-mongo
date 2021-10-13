package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Persona struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	Nombre        string             `redis:"nombre" bson:"nombre,omitempty" json:"nombre,omitempty" `
	Apellido      string             ` redis:"apellido" bson:"apellido,omitempty" json:"apellido,omitempty"`
	DireccionCasa Direccion          `redis:"direccion_casa" bson:"direccion_casa" json:"direccionCasa"`
}

type Personas []Persona

// Primero va la estructura a la que hace referencia, dps el nombre, y por ultimo los parametros
func (p Persona) Saludar() {
	fmt.Print("Hola")
}
