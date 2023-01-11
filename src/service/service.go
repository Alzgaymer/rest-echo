package electronic

import (
	"context"
	"log"
	mongodb "server/src/db/mongo"
	"server/src/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	mongodb.Mongodb
	logger *log.Logger
	arr    []*model.Product
}

var service Service

func init() {
	service.arr = make([]*model.Product, 0)
}

func New() *Service {
	return &service
}

func (s Service) SetClinet(client *mongo.Client) {
	s.Mongodb.MongoClient = client
}
func (s Service) SetURI(str string) {
	var err error
	s.MongoClient, err = mongo.NewClient(options.Client().ApplyURI(str))

	if err != nil {
		s.logger.Fatal(err)
	}
}
func (s Service) Connect(ctx context.Context) {
	err := s.MongoClient.Connect(ctx)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s Service) GetCollection(name string, options options.CollectionOptions) mongo.Collection {
	if s.Collection == nil {
		s.Collection = s.DB.Collection(name, &options)
	}
	return *s.Collection
}
