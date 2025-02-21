package report

import (
	"ppm/libnvth/internal/report/dto"
)

// Interactor interactor
type Interactor interface {
	GenerateReport(req dto.ReqReportGen) dto.ResReportGen
}
