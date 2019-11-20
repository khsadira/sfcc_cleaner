package prom_camp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func querySfccJSON(host string, target string, endpoint string, token string, field string, search string, start int, count int) []byte {
	client := &http.Client{}

	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["%s"],"search_phrase":"%s"}},"select":"(**)","start":%d,"count":%d}`, field, search, start, count)
	jsonBody := []byte(jsBody)

	query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s", host, target, endpoint)
	req, err := http.NewRequest("POST", query, bytes.NewBuffer(jsonBody))

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Println(err)
		return []byte("")
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return []byte("")
	}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	return buf
}

func querySfccDELETE(host string, site string, endpoint string, id string, token string) error {
	client := &http.Client{}

	query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s/%s", host, site, endpoint, id)
	req, err := http.NewRequest("DELETE", query, nil)

	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return err
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
