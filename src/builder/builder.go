package builder

import (
	"context"
	"log"
	"server/config"
	mongodb "server/src/db/mongo"
	"server/src/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBuilder struct {
	Service *service.Service
	Config  *config.ConfigDatabase
}

func (m MongoBuilder) Build() {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.Config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	client.Connect(context.Background())
	m.Service.Db = mongodb.New(client)
}
