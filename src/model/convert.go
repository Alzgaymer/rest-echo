package model

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MgoToModel(mgo *MgoProduct) *Product {
	return &Product{
		ID:           mgo.ID,
		ProductName:  mgo.ProductName,
		CreationTime: mgo.CreationTime,
	}
}

func ModelToMgo(mgo *Product) *MgoProduct {
	return &MgoProduct{
		ID:           mgo.ID,
		ProductName:  mgo.ProductName,
		CreationTime: mgo.CreationTime,
	}
}

func MgoToBsonD(mgo *MgoProduct) *bson.D {
	return &bson.D{
		{Key: "_id", Value: mgo.ID},
		{Key: "product_name", Value: mgo.ProductName},
		{Key: "creation_time", Value: mgo.CreationTime},
	}
}

func Decode(res *mongo.SingleResult) *Product {
	var (
		bsonproduct bson.M
		mgoproduct  MgoProduct
	)
	err := res.Decode(&bsonproduct)
	if err != nil {
		log.Print(err)
		return nil
	}
	bsonbytes, err := bson.Marshal(bsonproduct)
	if err != nil {
		log.Print(err)
		return nil
	}
	err = bson.Unmarshal(bsonbytes, &mgoproduct)
	if err != nil {
		return nil
	}
	return MgoToModel(&mgoproduct)
}
