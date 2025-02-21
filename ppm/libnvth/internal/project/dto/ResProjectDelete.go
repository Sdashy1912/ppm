package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResProjectDelete output for delete a vulnerability project usecase
type ResProjectDelete struct {
	Info basedto.ResInfo `json:"info"`
}
