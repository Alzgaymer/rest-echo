package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	_default_db         = "rest-echo-mongo"
	_default_collection = "item-collction"
)

type DB interface {
	New(client *mongo.Client) *Database
	Col() *mongo.Collection
}
type Database struct {
	db *mongo.Database
}

func New(client *mongo.Client) *Database {
	return &Database{db: client.Database(_default_db)}
}

func (d *Database) Col() *mongo.Collection {
	return d.db.Collection(_default_collection)
}
