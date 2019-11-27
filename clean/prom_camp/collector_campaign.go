package prom_camp

import (
	"encoding/json"
)

func getCampJSONs(host string, target string, token string, searchs []string) campData {
	var data campData

	endpoint := "/campaign_search"
	field := "enabled"
	for _, search := range searchs {
		var scan Scan
		var tmp campData

		count := 200

		bufTrue := querySfccJSON(host, target, endpoint, token, field, search, 0, count)
		json.Unmarshal(bufTrue, &scan)
		json.Unmarshal(bufTrue, &tmp)
		data.Hits = append(data.Hits, tmp.Hits...)

		for start := count; scan.Total > start; start += count {
			bufTrue = querySfccJSON(host, target, endpoint, token, field, search, start, count)
			json.Unmarshal(bufTrue, &tmp)
			data.Hits = append(data.Hits, tmp.Hits...)
		}
	}
	return data
}
