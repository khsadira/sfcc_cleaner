package utils

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendDeleteDataModule(w http.ResponseWriter, r *http.Request) {
	encodedBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%v\n", encodedBody)
	tmp := string(encodedBody)

	var bBody []byte

	if len(tmp) > 5 {
		bBody, _ = hex.DecodeString(tmp[5:])
	}
	body := string(bBody)

	splits := strings.Split(body, "&")
	print := fmt.Sprintf(`<html><p>ENDPAGE (BACK LOG TO SEE IF DB IS NORMALY SET</p><p>`)
	for i, split := range splits {
		splits[i] = strings.TrimSuffix(split, "=on")
		ret := strings.Split(splits[i], "*")

		if len(ret) == 5 {
			host := ret[0]
			site := ret[1]
			opts := ret[2]
			bID, _ := hex.DecodeString(ret[3])
			id := string(bID)
			endpoint := ret[4]
			print += fmt.Sprintf(`{host=%s, site=%s, opts=%s} <B>%s</B><br />`, host, site, opts, id)
			query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s/%s", host, site, endpoint, id)
			println(query)
			//	token, _ := GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
			//	QuerySfcc("GET", query, "Bearer", token, nil)
			println(host, site, opts, id)
		}
	}
	print += fmt.Sprintf(`</p>`)
	w.Write([]byte(print))
}
