package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ScanTotal struct {
	Total int `json:"total"`
}

func GetDataSFCC(query string, auth string) []interface{} {
	var ret []interface{}

	var data interface{}
	token, err := GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println("token was incorrect")
	}
	tmpQuery := query + "?start=0&count=200"
	buf := QuerySfcc("GET", tmpQuery, auth, token, nil)

	var scan ScanTotal
	json.Unmarshal(buf, &scan)
	json.Unmarshal(buf, &data)
	ret = append(ret, data)

	count := 200
	for start := count; scan.Total > start; start += count {
		tmpQuery = fmt.Sprintf("%s?start=%d&count=%d", query, start, count)
		buf = QuerySfcc("GET", tmpQuery, auth, token, nil)
		json.Unmarshal(buf, &data)
		ret = append(ret, data)
	}
	return ret
}

func QuerySfcc(method string, query string, auth string, token string, body []byte) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(method, query, bytes.NewBuffer(body))

	if err != nil {
		log.Println(err)
		return []byte("")
	}

	if auth == "Bearer" {
		req.Header.Add("Authorization", "Bearer "+token)
		if body != nil {
			req.Header.Add("Content-Type", "application/json")
		}
	} else if auth == "Basic" {
		req.Header.Add("Authorization", "Basic "+token)
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
