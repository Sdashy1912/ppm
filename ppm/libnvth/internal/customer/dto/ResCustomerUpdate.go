package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResCustomerUpdate output for update a customer usecase
type ResCustomerUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
