package semgrep

import (
	"encoding/json"
	"log"

	"github.com/parsiya/semgrep_go/run"
)

type ScanRequest struct {
	ScanID     string `json:"scan_id"`
	Repository string `json:"repository"`
	Path       string `json:"path"`
}

type ScanResult struct {
	Results []Results `json:"results"`
}

type Results struct {
	CheckID string `json:"check_id"`
	Message string `json:"message"`
}

func Scan(request ScanRequest) (ScanResult, error) {
	var scanResult ScanResult
	opts := run.Options{
		Output:    run.JSON,
		Paths:     []string{request.Path},
		Rules:     []string{"p/default"},
		Verbosity: run.Debug,
		Extra:     []string{"--no-rewrite-rule-ids"},
	}

	runResult, err := opts.Run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(runResult, &scanResult)
	if err != nil {
		log.Fatalf("error: %v", err)
		return scanResult, err
	}

	for _, result := range scanResult.Results {
		scanResult.Results = append(scanResult.Results, Results{
			CheckID: result.CheckID,
			Message: result.Message,
		})

	}
	return scanResult, nil
}
