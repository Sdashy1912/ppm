package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulTemplateCreate input for vul template create usecase
type ReqVulTemplateCreate struct {
	Info            basedto.ReqInfo `json:"-"`
	ID              string          `json:"id"`
	CategoryID      string          `json:"category_id"`
	Rating          string          `json:"rating"`
	Name            string          `json:"name"`
	IdentifierName  string          `json:"identifier_name"`
	Overview        string          `json:"overview"`
	DetectionMethod string          `json:"detection_method"`
	Description     string          `json:"description"`
	Condition       string          `json:"condition"`
	PossibleImpact  string          `json:"possible_impact"`
	Countermeasure  string          `json:"countermeasure"`
	Remarks         string          `json:"remarks"`
}
