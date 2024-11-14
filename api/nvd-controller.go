package api

import (
	"net/http"
	"snowden/pkg/nvd"
)

func ReadVulnerabilityByCve(w http.ResponseWriter, r *http.Request) error {
	cveId := r.URL.Query().Get("cveId")

	vulnerability, err := nvd.GetNvdModelByCveId(cveId)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	return WriteJson(w, http.StatusOK, vulnerability)
}

func ReadVulnerabilityByCwe(w http.ResponseWriter, r *http.Request) error {
	cweId := r.URL.Query().Get("cweId")

	vulnerabilities, err := nvd.GetNvdModelByCweId(cweId)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	return WriteJson(w, http.StatusOK, vulnerabilities)
}
