package pkg

import (
	"log"
	"snowden/client"
	"snowden/config"
)

var redisCache = config.NewRedisCache("localhost:6379", 0, 1)

var cweIds = []string{
	"CWE-284", "CWE-285", "CWE-614", "CWE-732", "CWE-922", "CWE-939", "CWE-275", "CWE-425", "CWE-319", "CWE-310", "CWE-312", "CWE-326", "CWE-327", "CWE-759", "CWE-760", "CWE-20", "CWE-74", "CWE-89", "CWE-77", "CWE-91", "CWE-564", "CWE-706", "CWE-116", "CWE-209", "CWE-602", "CWE-657", "CWE-841", "CWE-1037", "CWE-16", "CWE-297", "CWE-611", "CWE-916", "CWE-918", "CWE-920", "CWE-1104", "CWE-937", "CWE-477", "CWE-287", "CWE-798", "CWE-522", "CWE-620", "CWE-640", "CWE-352", "CWE-384", "CWE-494", "CWE-829", "CWE-353", "CWE-778", "CWE-223", "CWE-532", "CWE-117", "CWE-915",
}

func SeedCache() {
	for _, cweId := range cweIds {
		vuln, err := client.GetVulnerabilityByCweId(cweId)
		if err != nil {
			log.Fatal(err)
		}

		err = redisCache.SaveVulnerability(vuln.Cve)
		if err != nil {
			log.Fatal(err)
		}
	}
}
