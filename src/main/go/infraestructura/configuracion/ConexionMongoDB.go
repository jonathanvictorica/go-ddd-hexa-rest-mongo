package configuracion

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConectionMongoDB(collection string) *mongo.Collection {
	cargarConfiguraciones()

	usr := viper.GetString(`database.mongo.usr`)
	pwd := viper.GetString(`database.mongo.pwd`)
	host := viper.GetString(`database.mongo.host`)
	port := viper.GetInt(`database.mongo.port`)
	database := viper.GetString(`database.mongo.database`)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
