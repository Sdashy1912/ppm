package dao

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// IncidentDAOParam param for Create, Update
type IncidentDAOParam struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        `bson:"name"`
	CreatedAt      *time.Time    `bson:"created_at,omitempty"`
	LastUpdate     *time.Time    `bson:"last_update,omitempty"`
	Description    string        `bson:"description"`
	Solution       string        `bson:"solution"`
	Severity       int           `bson:"severity"`
	Status         string        `bson:"status"`
	Time           *time.Time    `bson:"time,omitempty"`
	Type           string        `bson:"type"`
	ProjectID      bson.ObjectId `bson:"project_id"`
	ResponsiblerID bson.ObjectId `bson:"responsibler_id"`
}
