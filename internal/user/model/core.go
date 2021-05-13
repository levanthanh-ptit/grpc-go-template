package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID primitive.ObjectID

var NilID ID = ID{}
