package dto

import (
	"time"

	"ppm/libnvth/internal/basedto"
)

// ReqProjectCreate input for vul project create usecase
type ReqProjectCreate struct {
	Info       basedto.ReqInfo `json:"-"`
	Name       string          `json:"name"`
	CustomerID string          `json:"customer_id"`
	StartDate  *time.Time      `json:"start_date,omitempty"`
	EndDate    *time.Time      `json:"end_date,omitempty"`
	Status     string          `json:"status"`
}
