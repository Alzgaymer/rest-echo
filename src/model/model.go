package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID           primitive.ObjectID  `json:"id" bson:"_ID"`
	Product_name string              `json:"product_name" validate:"required,min=4" bson:"product_name"`
	CreationTime primitive.Timestamp `json:"creation_time" bson:"_creation_time"`
}
