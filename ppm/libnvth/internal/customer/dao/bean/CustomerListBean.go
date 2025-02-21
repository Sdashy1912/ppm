package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// CustomerListBean bean
type CustomerListBean struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Email       string        `bson:"email"`
	PhoneNumber string        `bson:"phone_number"`
	Address     string        `bson:"address"`
	Industry    string        `bson:"industry"`
}
