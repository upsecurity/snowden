package nvd

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
			Metrics    Metrics `json:"metrics"`
			References []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
			} `json:"references"`
		} `json:"cve"`
	} `json:"vulnerabilities"`
}

type CWENVDModel struct {
	ResultsPerPage  int    `json:"resultsPerPage"`
	StartIndex      int    `json:"startIndex"`
	TotalResults    int    `json:"totalResults"`
	Format          string `json:"format"`
	Version         string `json:"version"`
	Timestamp       string `json:"timestamp"`
	Vulnerabilities []struct {
		Cve CWE `json:"cve"`
	} `json:"vulnerabilities"`
}

type CWE struct {
	ID               string           `json:"id"`
	SourceIdentifier string           `json:"sourceIdentifier"`
	Published        string           `json:"published"`
	LastModified     string           `json:"lastModified"`
	VulnStatus       string           `json:"vulnStatus"`
	CVETags          []struct{}       `json:"cveTags"`
	Descriptions     []Descriptions   `json:"descriptions"`
	Metrics          Metrics          `json:"metrics"`
	Weaknesses       []Weaknesses     `json:"weaknesses"`
	Configurations   []Configurations `json:"configurations"`
	References       []struct {
		URL    string   `json:"url"`
		Source string   `json:"source"`
		Tags   []string `json:"tags"`
	} `json:"references"`
}

type Descriptions struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Weaknesses struct {
	Source       string         `json:"source"`
	Type         string         `json:"type"`
	Descriptions []Descriptions `json:"descriptions"`
}

type Metrics struct {
	CvssV31 []CvssV3 `json:"cvssMetricV31"`
	CvssV30 []CvssV3 `json:"cvssMetricV30"`
	CvssV2  []CvssV2 `json:"cvssMetricV2"`
}

type CvssV3 struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	CvssData struct {
		Version               string  `json:"version"`
		VectorString          string  `json:"vectorString"`
		AttackVector          string  `json:"attackVector"`
		AttackComplexity      string  `json:"attackComplexity"`
		PrivilegesRequired    string  `json:"privilegesRequired"`
		UserInteraction       string  `json:"userInteraction"`
		Scope                 string  `json:"scope"`
		ConfidentialityImpact string  `json:"confidentialityImpact"`
		IntegrityImpact       string  `json:"integrityImpact"`
		AvailabilityImpact    string  `json:"availabilityImpact"`
		BaseScore             float64 `json:"baseScore"`
		BaseSeverity          string  `json:"baseSeverity"`
	} `json:"cvssData"`
	ExploitabilityScore float64 `json:"exploitabilityScore"`
	ImpactScore         float64 `json:"impactScore"`
}

type CvssV2 struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	CvssData struct {
		Version               string  `json:"version"`
		VectorString          string  `json:"vectorString"`
		AccessVector          string  `json:"accessVector"`
		AccessComplexity      string  `json:"accessComplexity"`
		Authentication        string  `json:"authentication"`
		ConfidentialityImpact string  `json:"confidentialityImpact"`
		IntegrityImpact       string  `json:"integrityImpact"`
		AvailabilityImpact    string  `json:"availabilityImpact"`
		BaseScore             float64 `json:"baseScore"`
	}
	BaseSeverity            string  `json:"baseSeverity"`
	ExploitabilityScore     float64 `json:"exploitabilityScore"`
	ImpactScore             float64 `json:"impactScore"`
	AcInsufInfo             bool    `json:"acInsufInfo"`
	ObtainAllPrivilege      bool    `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool    `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool    `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool    `json:"userInteractionRequired"`
}

type Configurations struct {
	Config []Config
}

type Config struct {
	Nodes []struct {
		Operator string `json:"operator"`
		Negate   bool   `json:"negate"`
		CpeMatch []struct {
			Vulnerable          bool   `json:"vulnerable"`
			Criteria            string `json:"criteria"`
			VersionEndExcluding string `json:"versionEndExcluding"`
			MatchCriteriaId     string `json:"matchCriteriaId"`
		} `json:"cpeMatch"`
	} `json:"nodes"`
}
