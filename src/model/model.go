package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `json:"id"`
	ProductName  string             `json:"product_name" validate:"required,min=4"`
	CreationTime time.Time          `json:"creation_time"`
}
