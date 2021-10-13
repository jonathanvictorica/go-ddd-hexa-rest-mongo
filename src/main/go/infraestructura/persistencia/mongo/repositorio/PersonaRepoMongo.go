package repositorio

import (
	"context"
	model2 "github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/model"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/dominio/repositorio"
	"github.com/JonathanMGuillermo/go-ddd-hexa-rest-mongo/src/main/go/infraestructura/configuracion"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type personaRepoMongo struct{}

var (
	collection *mongo.Collection
	ctx        context.Context
)

func NewPersonaRepoMongo() repositorio.PersonaRepositorio {
	collection = configuracion.NewConectionMongoDB("personas")
	ctx = context.Background()
	return &personaRepoMongo{}
}

func (*personaRepoMongo) CrearPersona(persona model2.Persona) error {
	var err error
	_, err = collection.InsertOne(ctx, persona)
	return err
}

func (*personaRepoMongo) ModificarPersona(persona model2.Persona) error {
	var err error
	old, _ := primitive.ObjectIDFromHex(persona.Id.String())
	filter := bson.M{"_id": old}
	update := bson.M{
		"$set": bson.M{
			"nombre":   persona.Nombre,
			"apellido": persona.Apellido,
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	return err

}

func (*personaRepoMongo) EliminarPersona(id string) error {
	var err error
	old, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": old}
	_, err = collection.DeleteOne(ctx, filter)
	return err
}

func (*personaRepoMongo) BuscarPersonaPorId(id string) (model2.Persona, error) {
	var persona model2.Persona
	old, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", old},
	}
	err := collection.FindOne(ctx, filter).Decode(&persona)
	return persona, err
}

func (*personaRepoMongo) BuscarPersonaPorNombre(nombre string) (model2.Persona, error) {
	var persona model2.Persona
	filter := bson.D{
		{"nombre", nombre},
	}
	err := collection.FindOne(ctx, filter).Decode(&persona)
	return persona, err
}

func (*personaRepoMongo) BuscarTodasPersonas() (model2.Personas, error) {
	filter := bson.D{}
	return convertirCursorPersonas(collection.Find(ctx, filter))
}

func convertirCursorPersonas(cur *mongo.Cursor, err error) (model2.Personas, error) {
	var personas model2.Personas
	if err != nil {
		return nil, err
	}
	// Recorremos el cursor
	for cur.Next(ctx) {
		var personaUnit model2.Persona
		err = cur.Decode(&personaUnit)
		if err != nil {
			return nil, err
		}
		personas = append(personas, personaUnit)
	}
	return personas, nil

}
