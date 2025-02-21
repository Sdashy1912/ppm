package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResCustomerGet output for get a customer usecase
type ResCustomerGet struct {
	Info     basedto.ResInfo    `json:"info"`
	Customer ResCustomerGetBean `json:"customer"`
}
