package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	NvdUrl = "https://services.nvd.nist.gov/rest/json/cves/2.0"
)

func GetVulnerabilityByCveId(cveId string) (Vulnerability, error) {
	var vulnerability CveVulnerability

	resp, err := http.Get(fmt.Sprintf("%s/?cveId=%s", NvdUrl, cveId))
	if err != nil {
		return Vulnerability{}, err
	}

	fmt.Println(resp.Body)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Vulnerability{}, err
	}

	err = json.Unmarshal(body, &vulnerability)
	if err != nil {
		return Vulnerability{}, err
	}

	vulnModel := modelVulnerability(vulnerability)

	return vulnModel, nil
}

func GetVulnerabilityByCweId(cweId string) (string, error) {
	var vulnerability string

	resp, err := http.Get(fmt.Sprintf("%s/?cweId=%s", NvdUrl, cweId))
	if err != nil {
		return vulnerability, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return vulnerability, err
	}

	err = json.Unmarshal(body, &vulnerability)
	if err != nil {
		return vulnerability, err
	}

	return vulnerability, nil
}

type CveVulnerability struct {
	ResultsPerPage  int    `json:"resultsPerPage"`
	StartIndex      int    `json:"startIndex"`
	TotalResults    int    `json:"totalResults"`
	Format          string `json:"format"`
	Version         string `json:"version"`
	Timestamp       string `json:"timestamp"`
	Vulnerabilities []struct {
		Cve struct {
			ID               string `json:"id"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Published        string `json:"published"`
			LastModified     string `json:"lastModified"`
			VulnStatus       string `json:"vulnStatus"`
			Descriptions     []struct {
				Lang  string `json:"lang"`
				Value string `json:"value"`
			} `json:"descriptions"`
			Metrics struct {
			} `json:"metrics"`
			References []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
			} `json:"references"`
		} `json:"cve"`
	} `json:"vulnerabilities"`
}

func modelVulnerability(vuln CveVulnerability) Vulnerability {
	var vulnerability Vulnerability

	vulnerability = vuln.Vulnerabilities[0]

	return vulnerability
}
