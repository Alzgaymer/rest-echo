package electronic

import (
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
