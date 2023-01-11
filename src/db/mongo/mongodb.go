package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongodb struct {
	MongoClient *mongo.Client
	DB          *mongo.Database
	Collection  *mongo.Collection
}
