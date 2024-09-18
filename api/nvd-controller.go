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
