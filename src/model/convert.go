package model

import "go.mongodb.org/mongo-driver/bson"

func MgoToModel(mgo *mgoProduct) *Product {
	return &Product{
		ID:           mgo.ID,
		Product_name: mgo.Product_name,
		CreationTime: mgo.CreationTime,
	}
}

func ModelToMgo(mgo *Product) *mgoProduct {
	return &mgoProduct{
		ID:           mgo.ID,
		Product_name: mgo.Product_name,
		CreationTime: mgo.CreationTime,
	}
}

func MgoToBson(mgo *mgoProduct) *bson.D {
	return &bson.D{
		{Key: "_id", Value: mgo.ID},
		{Key: "product_name", Value: mgo.Product_name},
		{Key: "creation_time", Value: mgo.CreationTime},
	}
}
