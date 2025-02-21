package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/incident/bean"
)

// ResIncidentList output for vulnerability incident list usecase
type ResIncidentList struct {
	Info      basedto.ResInfo         `json:"info"`
	Total     int                     `json:"total"`
	Incidents []bean.IncidentListBean `json:"incidents"`
}
