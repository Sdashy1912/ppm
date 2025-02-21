package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResTargetDetailAdd output for update target detail add usecase
type ResTargetDetailAdd struct {
	Info     basedto.ResInfo `json:"info"`
	Inserted int             `json:"inserted"`
}
