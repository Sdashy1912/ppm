package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// TargetDAOParam param for Create, Update
type TargetDAOParam struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Description string        `json:"description" bson:"description"`
	Platform    string        `json:"platform"`
	AssigneeID  bson.ObjectId `bson:"assignee_id"`
	ScopeID     bson.ObjectId `bson:"scope_id"`
}
