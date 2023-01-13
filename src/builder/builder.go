package builder

import (
	"context"
	"log"
	"server/config"
	service "server/src/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBuilder struct {
	Service *service.Service
	Config  *config.ConfigDatabase
}

func (m MongoBuilder) Build() {
	var err error
	m.Service.Mongodb.MongoClient, err = mongo.NewClient(
		options.Client().ApplyURI(m.Config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = m.Service.Mongodb.MongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	m.Service.Mongodb.DB = m.Service.Mongodb.MongoClient.Database("rest-echo")
	m.Service.Mongodb.Collection = m.Service.Mongodb.DB.Collection("item-collection")
	if err != nil {
		log.Fatal(err)
	}
}
