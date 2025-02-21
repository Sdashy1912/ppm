package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// UserBean bean
type UserBean struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Email       string        `json:"email" bson:"email"`
	PhoneNumber string        `json:"phone_number" bson:"phone_number"`
	IsActive    bool          `json:"is_active" bson:"is_active"`
	IsVerified  bool          `json:"is_verified" bson:"is_verified"`
	IsAdmin     bool          `json:"is_admin" bson:"is_admin"`
}
