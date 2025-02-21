package dto

import "ppm/libnvth/internal/basedto"

//ReqCustomerCreate input for create a new customer usecase
type ReqCustomerCreate struct {
	Info        basedto.ReqInfo `json:"-"`
	Name        string          `json:"name"`
	Address     string          `json:"address"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Industry    string          `json:"industry"`
}
