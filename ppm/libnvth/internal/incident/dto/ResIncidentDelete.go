package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResIncidentDelete output for delete a vulnerability incident usecase
type ResIncidentDelete struct {
	Info basedto.ResInfo `json:"info"`
}
