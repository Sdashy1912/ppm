package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResTargetDetailUpdate output for update target detail list usecase
type ResTargetDetailUpdate struct {
	Info     basedto.ResInfo `json:"info"`
	Inserted int             `json:"inserted"`
}
