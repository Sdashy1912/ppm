package dao

import (
	"ppm/libnvth/internal/project/bean"
	"gopkg.in/mgo.v2/bson"
)

// ProjectDAO interface
type ProjectDAO interface {
	List() ([]bean.ProjectBean, error)
	Get(ID bson.ObjectId) (bean.ProjectBean, error)
	Insert(param ProjectDAOParam) error
	Update(param ProjectDAOParam) error
	Remove(ID bson.ObjectId) error
	IsSafeToDelete(ID bson.ObjectId) (bool, error)
}
