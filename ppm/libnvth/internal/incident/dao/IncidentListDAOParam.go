package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// IncidentListDAOParam param for listing
type IncidentListDAOParam struct {
	ProjectID *bson.ObjectId
}
