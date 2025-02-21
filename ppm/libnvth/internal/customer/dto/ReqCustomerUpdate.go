package dto

import "ppm/libnvth/internal/basedto"

// ReqCustomerUpdate input for update a customer usecase
type ReqCustomerUpdate struct {
	Info        basedto.ReqInfo `json:"-"`
	ID          string          `json:"-"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Address     string          `json:"address"`
	Industry    string          `json:"industry"`
}
