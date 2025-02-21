package dao

import (
	"ppm/libnvth/internal/scope/bean"
	"gopkg.in/mgo.v2/bson"
)

// ScopeDAO interface
type ScopeDAO interface {
	Get(ID bson.ObjectId) (bean.ScopeBean, error)
	Insert(param ScopeDAOParam) error
	Update(param ScopeDAOParam) error
	Remove(ID bson.ObjectId) error
	IsSafeToDelete(ID bson.ObjectId) (bool, error)
}
