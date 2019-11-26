package blacklist

import (
	"github.com/khsadira/cleaner/auth"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func BlacklistDelModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {
		key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
		w.Header().Set("Authorization", key)
		t := template.New("example")
		t = template.Must(t.ParseFiles("template/get_bl_del.html"))
		err := t.ExecuteTemplate(w, "main", bList)

		if err != nil {
			log.Println(err)
		}
	}
}

func BlacklistDelValuesModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {

		body, _ := ioutil.ReadAll(r.Body)
		tab := strings.Split(string(body), "&")
		for _, str := range tab {
			str = strings.TrimSuffix(str, "=on")
			if strings.HasPrefix(str, "prom") {
				id := strings.TrimPrefix(str, "prom_")
				delBl(id, "prom")
			} else if strings.HasPrefix(str, "camp") {
				id := strings.TrimPrefix(str, "camp_")
				delBl(id, "camp")
			} else if strings.HasPrefix(str, "cstmgrp") {
				id := strings.TrimPrefix(str, "cstmgrp_")
				delBl(id, "cstmgrp")
			}
		}
		http.Redirect(w, r, "/blacklist/", 302)
	}
}

func removeIndexStock(s []Stock, index int) []Stock {
	return append(s[:index], s[index+1:]...)
}

func delBl(id string, opt string) {
	if opt == "prom" {
		for i, prom := range bList.Promotions {
			if prom.Encode == id {
				bList.Promotions = removeIndexStock(bList.Promotions, i)
			}
		}
	} else if opt == "camp" {
		for i, camp := range bList.Campaigns {
			if camp.Encode == id {
				bList.Campaigns = removeIndexStock(bList.Campaigns, i)
			}
		}
	} else if opt == "cstmgrp" {
		for i, cstmgrp := range bList.CustGrps {
			if cstmgrp.Encode == id {
				bList.CustGrps = removeIndexStock(bList.CustGrps, i)
			}
		}
	}
}
