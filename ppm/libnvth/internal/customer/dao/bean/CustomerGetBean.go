package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// CustomerGetBean bean
type CustomerGetBean struct {
	ID          bson.ObjectId `bson:"_id"`
	Email       string        `bson:"email"`
	Name        string        `bson:"name"`
	PhoneNumber string        `bson:"phone_number"`
	Industry    string        `bson:"industry"`
	Address     string        `bson:"address"`
	Projects    []ProjectBean `bson:"projects"`
}
