package model

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
