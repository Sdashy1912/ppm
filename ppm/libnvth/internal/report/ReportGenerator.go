package report

import (
	"ppm/libnvth/internal/report/model"
)

// Generator interface
type Generator interface {
	Generate(r model.Report) interface{}
}
