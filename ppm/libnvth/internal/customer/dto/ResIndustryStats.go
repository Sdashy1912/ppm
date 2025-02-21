package dto

import "ppm/libnvth/internal/basedto"

// ResIndustryStats response for customer stats
type ResIndustryStats struct {
	Info basedto.ResInfo        `json:"info"`
	Data []ResIndustryStatsData `json:"data"`
}
