package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	_default_db         = "rest-echo"
	_default_collection = "item-collection"
)

type Database struct {
	Db *mongo.Database
}

func New(client *mongo.Client) *mongo.Database {
	return client.Database(_default_db)
}

func (d *Database) Col() *mongo.Collection {
	return d.Db.Collection(_default_collection)
}
