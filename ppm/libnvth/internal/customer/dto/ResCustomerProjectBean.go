package dto

import "time"

// ResCustomerProjectBean bean
type ResCustomerProjectBean struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	Status    string     `json:"status"`
}
