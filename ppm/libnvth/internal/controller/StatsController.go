package controller

import (
	"net/http"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/stats"
	"github.com/go-chi/render"
)

// StatsController controller
type StatsController struct {
	interactor stats.Interactor
}

// NewStatsController initial
func NewStatsController(session *database.DBSession) StatsController {
	dao := stats.NewDAO(session)
	interactor := stats.NewInteractor(dao)
	return StatsController{interactor}
}

// Stats stats
func (controller StatsController) Stats(w http.ResponseWriter, r *http.Request) {
	result, err := controller.interactor.Stats()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	resp := make(map[string]interface{})
	resp["status_code"] = http.StatusOK
	resp["message"] = "OK"
	resp["data"] = result
	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
