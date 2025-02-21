package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqProjectList input for vul project list usecase
type ReqProjectList struct {
	Info basedto.ReqInfo `json:"-"`
}
