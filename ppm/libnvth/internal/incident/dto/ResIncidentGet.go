package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/incident/bean"
)

// ResIncidentGet output for get vulnerability incident output
type ResIncidentGet struct {
	Info     basedto.ResInfo   `json:"info"`
	Incident bean.IncidentBean `json:"incident"`
}
