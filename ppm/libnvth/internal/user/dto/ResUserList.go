package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/user/bean"
)

// ResUserList output for vulnerability user list usecase
type ResUserList struct {
	Info  basedto.ResInfo `json:"info"`
	Total int             `json:"total"`
	Users []bean.UserBean `json:"users"`
}
