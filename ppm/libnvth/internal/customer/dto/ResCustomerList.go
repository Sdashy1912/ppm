package dto

import "ppm/libnvth/internal/basedto"

// ResCustomerList output for list customer usecase
type ResCustomerList struct {
	Info      basedto.ResInfo       `json:"info"`
	Total     int                   `json:"total"`
	Customers []ResCustomerListBean `json:"customers"`
}
