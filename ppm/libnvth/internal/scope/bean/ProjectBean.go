package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// ProjectBean bean
type ProjectBean struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
