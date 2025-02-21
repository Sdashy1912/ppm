package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqIncidentDelete input for vul incident get usecase
type ReqIncidentDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
