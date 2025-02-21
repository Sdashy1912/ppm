package report

import (
	"fmt"
	"regexp"
	"strings"

	"ppm/libnvth/internal/report/model"
)

const header = `*********************************************************************************
CyRadar Information Security Joint Stock
2020 CyRadar Information Security Joint Stock, All Rights Reserved.
*********************************************************************************

Web Application Penetration Testing Service
Brief Report

--------------------------------------------------
%s
CyRadar Information Security Joint Stock
--------------------------------------------------
`

// TxtReportGenerator generate report in text format
type TxtReportGenerator struct {
}

var _ Generator = (*TxtReportGenerator)(nil)

func indentBlock(text, indent string) string {
	if text == "" {
		return indent
	}
	if text[len(text)-1:] == "\n" {
		result := ""
		for _, j := range strings.Split(text[:len(text)-1], "\n") {
			result += indent + j + "\n"
		}
		return result
	}
	result := ""
	for _, j := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		result += indent + j + "\n"
	}
	return result[:len(result)-1]
}

func (generator TxtReportGenerator) body(vulnerabilities []model.Vulnerability) string {
	vulList := "List of vulnerabilities\n\n"
	vulDetail := ""
	for idx, vul := range vulnerabilities {
		vulList += fmt.Sprintf("%-75sRisk:%s\n", fmt.Sprintf("■%d.%s", idx+1, vul.Name), vul.Rating)
		vulDetail += fmt.Sprintf("Number: %d\n", idx+1)
		vulDetail += fmt.Sprintf("■%d.%s\n\n", idx+1, vul.Name)
		vulDetail += "Risk Rating\n"
		vulDetail += indentBlock(vul.Rating, "  ") + "\n\n"
		vulDetail += "Possible Impact\n"
		vulDetail += indentBlock(vul.PossibleImpact, "  ") + "\n\n"
		vulDetail += "Description\n"
		description := vul.Description
		var re = regexp.MustCompile(`(?s)bq\.(.*?)bq\.`)
		description = re.ReplaceAllString(description, `--------------------------< START >-------------------------$1--------------------------<  END  >-------------------------`)
		vulDetail += indentBlock(description, "  ") + "\n\n"
		vulDetail += indentBlock("▼Target", "  ") + "\n"
		for _, pfTarget := range vul.PfTargets {
			//vulDetail += indentBlock("<pftarget>", "  ") + "\n"
			vulDetail += indentBlock(pfTarget.FuncName, "  ") + "\n"
			vulDetail += indentBlock(pfTarget.URL, "  ") + "\n"
			vulDetail += indentBlock(strings.Join(pfTarget.Parameters, ", "), "  ") + "\n\n"
			//vulDetail += indentBlock("</pftarget>", "  ") + "\n"
		}
		vulDetail += "Countermeasure\n"
		vulDetail += indentBlock(vul.Countermeasure, "  ") + "\n\n"
		vulDetail += "Remarks\n"
		vulDetail += indentBlock(vul.Remarks, "  ") + "\n\n"
		vulDetail += "--------------------------------------------------------------------------\n"
	}
	vulList += "\n--------------------------------------------------------------------------\n"
	return vulList + vulDetail
}

// Generate generate report in txt format
func (generator TxtReportGenerator) Generate(r model.Report) interface{} {
	output := fmt.Sprintf(header, r.CreatedAt.Format("January 02, 2006")) + "\n"
	output += generator.body(r.Target.Vulnerabilities) + "\n"
	return output
}
