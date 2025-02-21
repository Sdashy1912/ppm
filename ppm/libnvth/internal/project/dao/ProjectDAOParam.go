package dao

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ProjectDAOParam param for Create, Update
type ProjectDAOParam struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	CustomerID bson.ObjectId `bson:"customer_id"`
	StartDate  *time.Time    `bson:"start_date,omitempty"`
	EndDate    *time.Time    `bson:"end_date,omitempty"`
	Status     string        `bson:"status"`
}
