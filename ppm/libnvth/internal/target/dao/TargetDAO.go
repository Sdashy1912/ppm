package dao

import (
	"ppm/libnvth/internal/target/bean"
	"gopkg.in/mgo.v2/bson"
)

// TargetDAO interface
type TargetDAO interface {
	Get(ID bson.ObjectId) (bean.TargetBean, error)
	Insert(param TargetDAOParam) error
	Update(param TargetDAOParam) error
	Remove(ID bson.ObjectId) error
	GetDetailList(ID bson.ObjectId) ([]bean.TargetDetailBean, error)
	AddToDetailList(ID bson.ObjectId, detail bean.TargetDetailBean) error
	UpdateDetailList(ID bson.ObjectId, details []bean.TargetDetailBean) error
}
