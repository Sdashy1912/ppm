package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// CustomerBean bean
type CustomerBean struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
