package dto

import "ppm/libnvth/internal/basedto"

// ReqCustomerList input for list customers usecase
type ReqCustomerList struct {
	Info    basedto.ReqInfo
	OrderBy string
}
