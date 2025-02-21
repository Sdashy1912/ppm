package dto

import (
	"time"

	"ppm/libnvth/internal/basedto"
)

// ReqIncidentCreate input for vul incident create usecase
type ReqIncidentCreate struct {
	Info           basedto.ReqInfo `json:"-"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Solution       string          `json:"solution"`
	Severity       int             `json:"severity"`
	Status         string          `json:"status"`
	Time           *time.Time      `json:"time,omitempty"`
	Type           string          `json:"type"`
	ProjectID      string          `json:"project_id"`
	ResponsiblerID string          `json:"responsibler_id"`
}
