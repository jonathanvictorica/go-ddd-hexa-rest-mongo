package model

type Direccion struct {
	NombreCalle  string `redis:"nombreCalle" bson:"nombreCalle,omitempty" json:"nombreCalle,omitempty"`
	Altura       int    `redis:"altura" bson:"altura,omitempty" json:"altura,omitempty"`
	Piso         int    `redis:"piso" bson:"piso,omitempty" json:"piso,omitempty"`
	Departamento int    `redis:"departamento" bson:"departamento,omitempty" json:"departamento,omitempty"`
}
