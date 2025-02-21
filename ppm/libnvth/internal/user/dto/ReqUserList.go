package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqUserList input for vul user list usecase
type ReqUserList struct {
	Info basedto.ReqInfo `json:"-"`
}
