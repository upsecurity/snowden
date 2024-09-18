package nvd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snowden/config"
)

func GetNvdModelByCveId(cveId string) (NvdModel, error) {
	var vulnerability CveNvdModel
	nvdUrl := config.GetEnv("NVD_URL")
	resp, err := http.Get(fmt.Sprintf("%s?cveId=%s", nvdUrl, cveId))

	if err != nil {
		return NvdModel{}, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NvdModel{}, err
	}

	err = json.Unmarshal(body, &vulnerability)
	if err != nil {
		return NvdModel{}, err
	}

	vulnModel := modelNvdModel(vulnerability)

	return vulnModel, nil
}

type CveNvdModel struct {
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

func modelNvdModel(vuln CveNvdModel) NvdModel {
	vulnerability := vuln.Vulnerabilities[0]

	return vulnerability
}
