package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResTargetCreate output for create a new vulnerability scope usecase
type ResTargetCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
