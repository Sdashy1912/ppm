package dto

import "ppm/libnvth/internal/basedto"

// ResCustomerCreate output for create a new customer usecase
type ResCustomerCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
