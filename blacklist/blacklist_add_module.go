package blacklist

import (
	"encoding/hex"
	"github.com/khsadira/cleaner/auth"
	"net/http"
	"strings"
)

func BlacklistAddModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {
		key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
		w.Header().Set("Authorization", key)
		w.Write([]byte(`<html>
<p>If you want multi id, you can separate them by "," (don't put any space between ","'</p>
<form action="/blacklist/add/values" method="get" id="form1">
  promotion: <input type="text" name="prom"><br>
  campaigns: <input type="text" name="camp"><br>
  customer groups: <input type="text" name="cstm_grp"><br>
</form>

<button type="submit" form="form1" value="Submit">Submit</button>
</html>`))
	}
}

func BlacklistAddValuesModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {

		promQ := r.URL.Query()["prom"]
		campQ := r.URL.Query()["camp"]
		cstgrpQ := r.URL.Query()["cstm_grp"]

		for _, prom := range promQ {
			ids := strings.Split(prom, ",")
			for _, id := range ids {
				if len(id) > 0 {
					tmp := Stock{id, hex.EncodeToString([]byte(id))}
					bList.Promotions = append(bList.Promotions, tmp)
				}
			}
		}
		for _, camp := range campQ {
			ids := strings.Split(camp, ",")
			for _, id := range ids {
				if len(id) > 0 {
					tmp := Stock{id, hex.EncodeToString([]byte(id))}
					bList.Campaigns = append(bList.Campaigns, tmp)
				}
			}
		}
		for _, cstgrp := range cstgrpQ {
			ids := strings.Split(cstgrp, ",")
			for _, id := range ids {
				if len(id) > 0 {
					tmp := Stock{id, hex.EncodeToString([]byte(id))}
					bList.CustGrps = append(bList.CustGrps, tmp)
				}
			}
		}
		key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
		w.Header().Set("Authorization", key)
		http.Redirect(w, r, "/blacklist/", 302)
	}
}
