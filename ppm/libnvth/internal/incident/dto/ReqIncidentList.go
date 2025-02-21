package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqIncidentList input for vul incident list usecase
type ReqIncidentList struct {
	Info      basedto.ReqInfo `json:"-"`
	ProjectID string          `json:"-"`
}
