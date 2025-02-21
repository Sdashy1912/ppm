package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// ScopeBean bean
type ScopeBean struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Project ProjectBean   `json:"project" bson:"project"`
	Targets []TargetBean  `json:"targets,omitempty" bson:"targets,omitempty"`
}
