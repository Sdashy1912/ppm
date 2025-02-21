package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqUserUpdate input for vul user update usecase
type ReqUserUpdate struct {
	Info        basedto.ReqInfo `json:"-"`
	ID          string          `json:"-"`
	Name        string          `json:"name"`
	PhoneNumber string          `json:"phone_number"`
}
