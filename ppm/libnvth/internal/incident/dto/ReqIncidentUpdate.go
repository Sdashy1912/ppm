package dto

import (
	"time"

	"ppm/libnvth/internal/basedto"
)

// ReqIncidentUpdate input for vul incident update usecase
type ReqIncidentUpdate struct {
	Info           basedto.ReqInfo `json:"-"`
	ID             string          `json:"-"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Solution       string          `json:"solution"`
	Severity       int             `json:"severity"`
	Status         string          `json:"status"`
	Time           *time.Time      `json:"time,omitempty"`
	Type           string          `json:"type"`
	ResponsiblerID string          `json:"responsibler_id"`
}
