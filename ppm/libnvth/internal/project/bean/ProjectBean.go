package bean

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ProjectBean bean
type ProjectBean struct {
	ID        bson.ObjectId  `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	StartDate *time.Time     `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time     `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Status    string         `json:"status" bson:"status"`
	Customer  CustomerBean   `json:"customer,omitempty" bson:"customer,omitempty"`
	Scopes    []ScopeBean    `json:"scopes,omitempty" bson:"scopes,omitempty"`
	Incidents []IncidentBean `json:"incidents" bson:"incidents"`
}
