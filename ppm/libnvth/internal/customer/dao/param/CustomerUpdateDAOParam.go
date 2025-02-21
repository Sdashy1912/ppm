package param

import (
	"gopkg.in/mgo.v2/bson"
)

// CustomerUpdateDAOParam params for create a new customer
type CustomerUpdateDAOParam struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Address     string        `bson:"address"`
	Email       string        `bson:"email"`
	PhoneNumber string        `bson:"phone_number"`
	Industry    string        `bson:"industry"`
}
