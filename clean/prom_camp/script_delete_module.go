package prom_camp

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ScriptDeleteModule(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		str := formDelete(body)
		w.Write([]byte(str))
	}
}

func formDelete(bBody []byte) string {
	var p, c int

	body := string(bBody)

	splits := strings.Split(body, "&")
	str := `<html><p>You are going to delete all these campaigns and promotions</p><p>`
	for i, split := range splits {
		splits[i] = strings.TrimSuffix(split, "=on")
		ret := strings.Split(splits[i], "*")

		if len(ret) == 4 {

			host := ret[0]
			site := ret[1]
			opts := ret[2]
			bID, _ := hex.DecodeString(ret[3])
			id := string(bID)
			println("delete:", host+":"+site+":"+opts+":"+id)
			str += fmt.Sprintf(`{host=%s, site=%s, opts=%s} <B>%s</B><br />`, host, site, opts, id)
			if opts[:4] == "prom" {
				p++
			} else {
				c++
			}

		}
	}
	str += fmt.Sprintf(`</p><p>Total Promotions: %d<br />Total Campaigns: %d<br /><br />Total: %d</p>`, p, c, p+c)
	cBody := hex.EncodeToString(bBody)
	str += fmt.Sprintf(`<form action="/promo/clean/del" method="post">
    <button name="data" value="%s">Confirmed</button>
</form>`, cBody)
	return str
}
