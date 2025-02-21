package controller

import (
    "net/http"

    "ppm/libnvth/internal/basedto"

    "ppm/libnvth/internal/database"
    "ppm/libnvth/internal/report"
    "ppm/libnvth/internal/report/dao"
    "ppm/libnvth/internal/report/dto"
    "github.com/go-chi/render"
)

// ReportController controller
type ReportController struct {
    interactor report.Interactor
}

// NewReportController initial
func NewReportController(session *database.DBSession) ReportController {
    dao := dao.NewReportDAOImpl(session)
    interactor := report.NewInteractorImpl(dao)
    return ReportController{interactor}
}

// Generate generate reports for a target
func (controller ReportController) Generate(w http.ResponseWriter, r *http.Request) {
    req := dto.ReqReportGen{}
    if err := render.DecodeJSON(r.Body, &req); err != nil {
    	render.Status(r, http.StatusBadRequest)
    	render.JSON(w, r, err.Error())
    	return
    }
    resp := controller.interactor.GenerateReport(req)
    if resp.Info.StatusCode != basedto.StatusCodeOK {
        render.Status(r, resp.Info.StatusCode)
        render.JSON(w, r, resp)
        return
    }
    render.PlainText(w, r, resp.Report.(string))
}
