package blacklist

import (
	"encoding/json"
	"github.com/khsadira/cleaner/auth"
	"net/http"
)

func BlacklistSendModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {
		if r.Method == "GET" {
			ret, err := json.Marshal(&bList)
			if err != nil {
				ret = []byte(`we got an error`)
			}
			w.Write(ret)
		}
	}
}
