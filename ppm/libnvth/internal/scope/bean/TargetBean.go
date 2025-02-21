package bean

import "gopkg.in/mgo.v2/bson"

// TargetBean bean
type TargetBean struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Platform    string        `json:"platform" bson:"platform"`
	Assignee    AssigneeBean  `json:"assignee" bson:"assignee"`
}
