package dto

import (
	"time"

	"ppm/libnvth/internal/basedto"
)

// ReqProjectUpdate input for vul project update usecase
type ReqProjectUpdate struct {
	Info       basedto.ReqInfo `json:"-"`
	ID         string          `json:"-"`
	Name       string          `json:"name"`
	CustomerID string          `json:"customer_id"`
	StartDate  *time.Time      `json:"start_date,omitempty"`
	EndDate    *time.Time      `json:"end_date,omitempty"`
	Status     string          `json:"status"`
}
