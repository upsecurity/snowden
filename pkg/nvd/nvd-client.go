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

func GetNvdModelByCweId(cweId string) (CWENVDModel, error) {
	var vuln CWENVDModel
	nvdUrl := config.GetEnv("NVD_URL")
	resp, err := http.Get(fmt.Sprintf("%s?cweId=%s", nvdUrl, cweId))

	if err != nil {
		return CWENVDModel{}, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CWENVDModel{}, err
	}

	err = json.Unmarshal(body, &vuln)
	if err != nil {
		return CWENVDModel{}, err
	}

	return vuln, nil
}

func modelNvdModel(vuln CveNvdModel) NvdModel {
	vulnerability := vuln.Vulnerabilities[0]

	return vulnerability
}
