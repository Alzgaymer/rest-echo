package model

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func Decode(res *mongo.SingleResult) *Product {
	var (
		err         error
		bsonproduct bson.M
		mgoproduct  mgoProduct
	)
	res.Decode(&bsonproduct)
	if err != nil {
		log.Print(err)
		return nil
	}
	bsonbytes, err := bson.Marshal(bsonproduct)
	if err != nil {
		log.Print(err)
		return nil
	}
	bson.Unmarshal(bsonbytes, &mgoproduct)
	return MgoToModel(&mgoproduct)
}
