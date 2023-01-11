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
	logger     *log.Logger
	Collection []*model.Product
}

var service Service

func init() {
	service.Collection = make([]*model.Product, 0)
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
func (s Service) Connect(ctx context.Context, dbName string) {
	err := s.MongoClient.Connect(ctx)
	if err != nil {
		s.logger.Fatal(err)
	}
	s.DB = s.MongoClient.Database(dbName)
}

func (s Service) GetCollection(name string) *mongo.Collection {
	if s.Collection == nil {
		s.Mongodb.Collection = s.DB.Collection(name)
	}
	return s.Mongodb.Collection
}

func (s Service) InsertMany(ctx context.Context, products ...interface{}) *mongo.InsertManyResult {
	result, err := s.Mongodb.Collection.InsertMany(ctx, products)
	if err != nil {
		s.logger.Fatal(err)
	}
	return result
}

func (s Service) InsertOne(ctx context.Context, product interface{}) *mongo.InsertOneResult {
	result, err := s.Mongodb.Collection.InsertOne(ctx, product)
	if err != nil {
		s.logger.Fatal(err)
	}
	return result
}
