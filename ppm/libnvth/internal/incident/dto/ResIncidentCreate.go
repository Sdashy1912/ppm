package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResIncidentCreate output for create a new vulnerability incident usecase
type ResIncidentCreate struct {
	Info       basedto.ResInfo `json:"info"`
	InsertedID string          `json:"inserted_id"`
}
