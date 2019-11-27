package utils

import (
	"encoding/json"
	"fmt"
)

func GetSites(host string) []string {
	token, _ := GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")

	sites := findSites(host, []string{}, token, 0, 200)
	return sites
}

func findSites(host string, sites []string, token string, start int, count int) []string {
	query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites?start=%d&count=%d", host, start, count)
	buf := QuerySfcc("GET", query, "Bearer", token, nil)

	var data Sites
	json.Unmarshal(buf, &data)

	for i := 0 ; i < data.Count; i++ {
		sites = append(sites, data.Data[i].ID)
	}
	if data.Total >= start+data.Count {
		sites = findSites(host, sites, token, start+count, count)
	}
	return sites

}