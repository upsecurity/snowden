package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"snowden/client/model"
)

func GetVulnerabilityByCweId(cweId string) (model.Vulnerability, error) {
	var vulnerability model.CompleteCwe
	nvdUrl := os.Getenv("NVD_URL")

	resp, err := http.Get(fmt.Sprintf("%s/?cweId=%s", nvdUrl, cweId))
	if err != nil {
		return model.Vulnerability{}, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Vulnerability{}, err
	}

	err = json.Unmarshal(body, &vulnerability)
	if err != nil {
		return model.Vulnerability{}, err
	}

	vulnModel := model.MarshallCweVulnerability(vulnerability)

	return vulnModel, nil
}

func GetVulnerabilityByCveId(cveId string) ([]model.Vulnerability, error) {
	var vulnerability model.CompleteCve
	nvdUrl := os.Getenv("NVD_URL")

	resp, err := http.Get(fmt.Sprintf("%s/?cveId=%s", nvdUrl, cveId))
	if err != nil {
		return []model.Vulnerability{}, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []model.Vulnerability{}, err
	}

	err = json.Unmarshal(body, &vulnerability)
	if err != nil {
		return []model.Vulnerability{}, err
	}

	vulnModel := model.MarshallCveVulnerability(vulnerability)

	return vulnModel, nil
}
