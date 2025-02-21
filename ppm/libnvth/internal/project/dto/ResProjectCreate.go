package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResProjectCreate output for create a new vulnerability project usecase
type ResProjectCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
