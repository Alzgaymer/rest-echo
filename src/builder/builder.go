package builder

import (
	"context"
	"server/config"
	mongodb "server/src/db/mongo"
	"server/src/service"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBuilder struct {
	Service *service.Service
	Config  *config.ConfigDatabase
	Log     echo.Logger
}

func (m MongoBuilder) Build() {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.Config.MongoURI))
	if err != nil {
		m.Log.Fatal(err)
	}
	err = client.Connect(context.Background())
	m.Log.Print(err)
	m.Service.Db = mongodb.New(client)
}
