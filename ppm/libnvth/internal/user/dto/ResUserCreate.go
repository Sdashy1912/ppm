package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResUserCreate output for create a new vulnerability user usecase
type ResUserCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
