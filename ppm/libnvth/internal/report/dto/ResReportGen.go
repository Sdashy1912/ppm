package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResReportGen response
type ResReportGen struct {
	Info   basedto.ResInfo `json:"info"`
	Report interface{}     `json:"report"`
}
