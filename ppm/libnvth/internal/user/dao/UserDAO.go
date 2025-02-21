package dao

import (
	"ppm/libnvth/internal/user/bean"
	"gopkg.in/mgo.v2/bson"
)

// UserDAO interface
type UserDAO interface {
	List() ([]bean.UserBean, error)
	Get(ID bson.ObjectId) (bean.UserBean, error)
	GetByEmail(email string) (bean.UserBean, error)
	GetByEmailAndPassword(email string, password string) (bean.UserBean, error)
	Insert(param UserCreateDAOParam) error
	Update(param UserUpdateDAOParam) error
	Remove(ID bson.ObjectId) error
	IsSafeToDelete(ID bson.ObjectId) (bool, error)
}
