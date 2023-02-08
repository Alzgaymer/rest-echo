package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MgoProduct struct {
	ID           primitive.ObjectID `bson:"_id"`
	Product_name string             `bson:"product_name"`
	CreationTime time.Time          `bson:"creation_time"`
}
