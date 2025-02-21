package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// AssigneeBean bean for user who is assigned to the scope
type AssigneeBean struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
