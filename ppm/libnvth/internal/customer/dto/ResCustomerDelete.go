package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResCustomerDelete output for delete a customer usecase
type ResCustomerDelete struct {
	Info basedto.ResInfo `json:"info"`
}
