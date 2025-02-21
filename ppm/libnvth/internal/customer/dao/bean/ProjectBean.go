package bean

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ProjectBean bean
type ProjectBean struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	StartTime *time.Time    `bson:"start_time,omitempty"`
	EndTime   *time.Time    `bson:"end_time,omitempty"`
	Status    string        `bson:"status"`
}
