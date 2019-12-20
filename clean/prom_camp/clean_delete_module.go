package prom_camp

import (
	"encoding/hex"
	"fmt"
	"github.com/khsadira/cleaner/clean/utils"
	"html/template"
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

		var p, c int

		splits := strings.Split(body, "&")
		str := `<html><p>ENDPAGE (BACK LOG TO SEE IF DB IS NORMALY SET'</p><p>`
		for i, split := range splits {
			splits[i] = strings.TrimSuffix(split, "=on")
			ret := strings.Split(splits[i], "*")

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
					println("BACK ON DELETE:", host+":"+site+":"+opts+":"+endpoint+":"+id)
					query := fmt.Sprintf("https://%s/s/-/dw/data/v19_8/sites/%s/%s/%s", host, site, endpoint, template.URLQueryEscaper(id))
					token, _ := utils.GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
					println(query, token)
					utils.QuerySfcc("DELETE", query, "Bearer", token, nil)
				}
				println("")
				str += fmt.Sprintf(`{host=%s, site=%s, opts=%s} <B>%s</B><br />`, host, site, opts, id)
				if opts[:4] == "prom" {
					p++
				} else {
					c++
				}

			}
		}
		str += fmt.Sprintf(`</p><p>Total Promotions: %d<br />Total Campaigns: %d<br /><br />Total: %d</p>`, p, c, p+c)
		w.Write([]byte(str))
	}
}
