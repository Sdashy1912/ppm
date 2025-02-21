package dao

import (
	"ppm/libnvth/internal/incident/bean"
	"gopkg.in/mgo.v2/bson"
)

// IncidentDAO interface
type IncidentDAO interface {
	List(IncidentListDAOParam) ([]bean.IncidentListBean, error)
	Get(ID bson.ObjectId) (bean.IncidentBean, error)
	Insert(param IncidentDAOParam) error
	Update(param IncidentDAOParam) error
	Remove(ID bson.ObjectId) error
}
