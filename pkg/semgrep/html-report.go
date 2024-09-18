package semgrep

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"snowden/pkg"
	"time"

	"github.com/parsiya/semgrep_go/output"
)

var tpl string

type Report struct {
	NumberOfResults int
	ByRuleID        []output.HitMapRow
	ByFilePath      []output.HitMapRow
}

func NewHTMLReport(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	output, err := output.Deserialize(data)
	if err != nil {
		return err
	}

	report := Report{
		NumberOfResults: len(data),
		ByRuleID:        output.RuleIDHitMap(true),
		ByFilePath:      output.FilePathHitMap(true),
	}

	temp, err := template.New("report").Parse(tpl)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	if err = temp.Execute(&buffer, report); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s-report.html", time.Now().Format(time.RFC1123))

	log.Printf("Writing report to %s", fileName)
	return pkg.WriteFile(fileName, buffer.Bytes())
}
