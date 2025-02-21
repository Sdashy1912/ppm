package model

import "time"

// Report report
type Report struct {
	CreatedAt  time.Time
	TemplateID string
	Target     Target
}
