package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/user/bean"
)

// ResUserGet output for get vulnerability user output
type ResUserGet struct {
	Info basedto.ResInfo `json:"info"`
	User bean.UserBean   `json:"user"`
}
