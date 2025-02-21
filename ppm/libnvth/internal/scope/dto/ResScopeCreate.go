package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResScopeCreate output for create a new vulnerability scope usecase
type ResScopeCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
