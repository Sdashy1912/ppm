package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqUserCreate input for vul user create usecase
type ReqUserCreate struct {
	Info        basedto.ReqInfo `json:"-"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
}
