package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// UserCreateDAOParam param for Create, Update
type UserCreateDAOParam struct {
	ID          bson.ObjectId `bson:"_id"`
	Email       string        `bson:"email"`
	Password    string        `bson:"password"`
	Name        string        `bson:"name"`
	PhoneNumber string        `bson:"phone_number"`
	IsActive    bool          `bson:"is_active"`
	IsVerified  bool          `bson:"is_verified"`
	IsAdmin     bool          `bson:"is_admin"`
}
