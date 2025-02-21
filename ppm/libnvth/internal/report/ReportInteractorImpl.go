package report

import (
    "time"

    "ppm/libnvth/internal/report/dao"
    "ppm/libnvth/internal/report/dto"
    "ppm/libnvth/internal/report/model"
    "gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
    dao       dao.ReportDAO
    generator Generator
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.ReportDAO) InteractorImpl {
    generator := TxtReportGenerator{}
    return InteractorImpl{dao, generator}
}

var _ Interactor = (*InteractorImpl)(nil)

// GenerateReport generate report
func (interactor InteractorImpl) GenerateReport(req dto.ReqReportGen) dto.ResReportGen {
    resp := dto.ResReportGen{}
    if !bson.IsObjectIdHex(req.TargetID) {
        resp.Info.SetStatusUnprocessableEntity(map[string]string{"target_id": "invalid target_id"})
        return resp
    }
    target, err := interactor.dao.GetTargetInfo(bson.ObjectIdHex(req.TargetID))
    if err != nil {
        if err.Error() == "not found" {
            resp.Info.SetStatusNotFound()
        } else {
            resp.Info.SetStatusInternalServerError()
        }
        return resp
    }
    var vulsToReport []model.Vulnerability
    now := time.Now()
    switch req.Mode {
    case 0:
        vulsToReport = target.Vulnerabilities
    case 1:
        for _, vul := range target.Vulnerabilities {
            if vul.LastUpdate == nil || !dateEquals(*vul.LastUpdate, now) {
                continue
            }
            vulsToReport = append(vulsToReport, vul)
        }
    case 2:
        for _, id := range req.VulIDs {
            vul, err := target.GetVulByID(id)
            if !err {
                resp.Info.SetStatusUnprocessableEntity(map[string]string{"vul_ids": "invalid vulnerability ids"})
                return resp
            }
            vulsToReport = append(vulsToReport, vul)
        }
    }
    target.Vulnerabilities = vulsToReport
    rp := model.Report{
        CreatedAt: time.Now(),
        Target:    target,
    }
    doc := interactor.generator.Generate(rp)
    resp.Info.SetStatusOK("ok")
    resp.Report = doc
    return resp
}

func dateEquals(date1, date2 time.Time) bool {
    y1, m1, d1 := date1.Date()
    y2, m2, d2 := date2.Date()
    return y1 == y2 && m1 == m2 && d1 == d2
}
