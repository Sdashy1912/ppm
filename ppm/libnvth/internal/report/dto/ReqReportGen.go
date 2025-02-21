package dto

import (
    "ppm/libnvth/internal/basedto"
)

// ReqReportGen request
type ReqReportGen struct {
    Info       basedto.ReqInfo `json:"-"`
    TargetID   string          `json:"target_id"`
    TemplateID string          `json:"template_id"`
    Mode       int             `json:"mode"`
    VulIDs     []string        `json:"vul_ids"`
}
