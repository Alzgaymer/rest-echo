package service

import (
	"context"
	"log"

	mongodb "server/src/db/mongo"
)

type Service struct {
	mongodb.Database
	logger *log.Logger
}

var service Service

func New() *Service {
	return &service
}

// func (s Service) InsertOne(ctx context.Context, obj interface{}) *mongo.InsertOneResult {

// }

func (s Service) Disconnect(ctx context.Context) {
	log.Printf("Disconnecting from mongo...")
	err := s.Db.Client().Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Disconnected succesfully")
}
