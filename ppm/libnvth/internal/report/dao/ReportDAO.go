package dao

import (
	"ppm/libnvth/internal/report/model"
	"gopkg.in/mgo.v2/bson"
)

// ReportDAO dao
type ReportDAO interface {
	GetTargetInfo(targetID bson.ObjectId) (model.Target, error)
}
