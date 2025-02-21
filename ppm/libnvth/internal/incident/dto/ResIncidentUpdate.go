package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResIncidentUpdate output for update a vulnerability incident usecase
type ResIncidentUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
