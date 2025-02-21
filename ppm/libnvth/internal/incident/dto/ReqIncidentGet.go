package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqIncidentGet input for vul incident get usecase
type ReqIncidentGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
