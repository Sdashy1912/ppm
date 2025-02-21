package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/report/model"
	"gopkg.in/mgo.v2/bson"
)

// ReportDAOImpl implementations
type ReportDAOImpl struct {
	session *database.DBSession
}

var _ ReportDAO = (*ReportDAOImpl)(nil)

// NewReportDAOImpl return a new ReportDAOImpl instance
func NewReportDAOImpl(session *database.DBSession) ReportDAOImpl {
	return ReportDAOImpl{session}
}

// GetTargetInfo collects all needed informations from target_id
func (dao ReportDAOImpl) GetTargetInfo(targetID bson.ObjectId) (model.Target, error) {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	target := model.Target{}
	err := collection.FindId(targetID).One(&target)
	return target, err
}
