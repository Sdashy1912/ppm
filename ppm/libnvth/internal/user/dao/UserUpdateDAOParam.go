package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// UserUpdateDAOParam param for update case
type UserUpdateDAOParam struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	PhoneNumber string        `bson:"phone_number"`
}
