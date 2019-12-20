package prom_camp

import (
	"encoding/hex"
	"fmt"
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func CleanDelModule(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		encodedBody, _ := ioutil.ReadAll(r.Body)
		tmp := string(encodedBody)
		decodedBody, _ := hex.DecodeString(tmp[5:])

		body := string(decodedBody)
		splits := strings.Split(body, "&")

		sizeSplits := len(splits)
		ch := make(chan bool, sizeSplits)

		for _, split := range splits {
			go scriptDeleteRoutine(split, ch)
		}
		for i := 0; i < sizeSplits; i++ {
			<-ch
		}
	}
	w.Write([]byte(`<!DOCTYPE HTML><html><p>ENDPAGE (BACK LOG TO SEE IF DB IS NORMALY SET</p></html>`))
}

func scriptDeleteRoutine(split string, ch chan bool) {
	split = strings.TrimSuffix(split, "=on")
	ret := strings.Split(split, "*")

	if len(ret) == 4 {
		var endpoint string

		host := ret[0]
		site := ret[1]
		opts := ret[2]
		bID, _ := hex.DecodeString(ret[3])
		id := string(bID)
		if host == "store-dev.ubi.com" { //temporary protection to only del on dev
			if opts[:4] == "prom" {
				endpoint = "promotions"
			} else {
				endpoint = "campaigns"
			}
			query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s/%s", host, site, endpoint, utils.ReworkID(id))
			println("Delete:", host, site, endpoint, id, "\n"+"Query:", query)
			token, _ := utils.GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
			utils.QuerySfcc("DELETE", query, "Bearer", token, nil)
		}
	}

	ch <- true
}
