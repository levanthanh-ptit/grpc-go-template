package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID primitive.ObjectID

func (id ID) ToString() string {
	var val interface{} = id
	return val.(primitive.ObjectID).String()
}

var NilID ID = ID{}
