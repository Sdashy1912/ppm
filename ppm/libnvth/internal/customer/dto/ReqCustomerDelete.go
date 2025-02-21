package dto

import "ppm/libnvth/internal/basedto"

// ReqCustomerDelete input for delete a customer usecase
type ReqCustomerDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
}
