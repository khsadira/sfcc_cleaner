package utils

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

func SendConfirmDelete(w http.ResponseWriter, bBody []byte, nameData string, path string) {
	var countData int

	body := string(bBody)

	splits := strings.Split(body, "&")
	print := fmt.Sprintf(`<html><p>You are going to delete all these %s</p><p>`, nameData)
	for i, split := range splits {
		splits[i] = strings.TrimSuffix(split, "=on")
		ret := strings.Split(splits[i], "*")

		if len(ret) == 4 {
			host := ret[0]
			site := ret[1]
			opts := ret[2]
			bID, _ := hex.DecodeString(ret[3])
			id := string(bID)
			print += fmt.Sprintf(`{host=%s, site=%s, opts=%s} <B>%s</B><br />`, host, site, opts, id)
			countData++
		}
	}
	print += fmt.Sprintf(`</p><p>Total %s: %d</p>`, nameData, countData)
	cBody := hex.EncodeToString(bBody)
	print += fmt.Sprintf(`<form action="%s" method="post">
    <button name="data" value="%s">Confirmed</button>
</form>`, path, cBody)
	w.Write([]byte(print))
}
