package bean

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// IncidentBean bean
type IncidentBean struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	CreatedAt    *time.Time    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	LastUpdate   *time.Time    `json:"last_update,omitempty" bson:"last_update,omitempty"`
	Description  string        `json:"description" bson:"description"`
	Solution     string        `json:"solution" bson:"solution"`
	Severity     int           `json:"severity" bson:"severity"`
	Status       string        `json:"status" bson:"status"`
	Time         *time.Time    `json:"time,omitempty" bson:"time,omitempty"`
	Type         string        `json:"type" bson:"type"`
	Project      ProjectBean   `json:"project" bson:"project"`
	Customer     CustomerBean  `json:"customer" bson:"customer"`
	Responsibler PentesterBean `json:"responsibler" bson:"responsibler"`
}
