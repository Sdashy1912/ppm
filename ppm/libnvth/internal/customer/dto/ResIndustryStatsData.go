package dto

// ResIndustryStatsData bean
type ResIndustryStatsData struct {
	Industry  string `json:"industry" bson:"industry"`
	Customers int    `json:"customers" bson:"customers"`
}
