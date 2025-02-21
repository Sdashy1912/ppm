package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// UserChangePasswordDAOParam param
type UserChangePasswordDAOParam struct {
	ID       bson.ObjectId `bson:"_id"`
	Password string        `bson:"password"`
}
