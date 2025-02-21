package dto

import "ppm/libnvth/internal/basedto"

// ReqCustomerGet input for get a customer usecase
type ReqCustomerGet struct {
	Info basedto.ReqInfo
	ID   string
}
