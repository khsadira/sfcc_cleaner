package utils

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
***	TO SEND DATA TO BE DELETED (GENERAL FUNCTION, CAN BE USED WITH ANY DATA)
 */

func SendDeleteDataModule(w http.ResponseWriter, r *http.Request) {
	encodedBody, _ := ioutil.ReadAll(r.Body)
	tmp := string(encodedBody)

	var bBody []byte

	if len(tmp) > 5 {
		bBody, _ = hex.DecodeString(tmp[5:])
	}

	body := string(bBody)
	splits := strings.Split(body, "&")

	sizeSplits := len(splits)
	ch := make(chan bool, sizeSplits)

	for _, split := range splits {
		go scriptDeleteRoutine(split, ch)

	}
	for i := 0; i < sizeSplits; i++ {
		<-ch
	}
	w.Write([]byte(`<!DOCTYPE HTML><html><p>ENDPAGE (BACK LOG TO SEE IF DB IS NORMALY SET</p></html>`))
	println("DELETE DONE")
}

func scriptDeleteRoutine(split string, ch chan bool) {
	split = strings.TrimSuffix(split, "=on")
	ret := strings.Split(split, "*")

	if len(ret) == 5 {
		host := ret[0]
		site := ret[1]
		bID, _ := hex.DecodeString(ret[3])
		id := string(bID)
		endpoint := ret[4]
		query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s/%s", host, site, endpoint, ReworkID(id))
		println("Delete:", host, site, endpoint, id, "\n"+"Query:", query)
		token, _ := GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
		QuerySfcc("DELETE", query, "Bearer", token, nil)
	}

	ch <- true
}
