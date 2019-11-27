package prom_camp

import (
	"bytes"
	"encoding/json"
	"github.com/khsadira/cleaner/auth"
	"github.com/khsadira/cleaner/clean/utils"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func CleanModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {

		if r.Method == "POST" {
			var send HostsScan

			fillSend(r, &send)
			ret, err := json.Marshal(send)

			if err != nil {
				log.Println(err)
			}
			resp, err := auth.SendPostAuth("http://localhost:9250/promo/script/info/", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC", bytes.NewBuffer(ret))
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
			generateDelForm(w, body)
		} else {
			generateInfoForm(w)
		}
	}
}

func generateDelForm(w http.ResponseWriter, body []byte) {
	var send []hostsStruct
	var bList BlackList

	resp, err := auth.SendGetAuth("http://localhost:9250/blacklist/send/", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	blBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(blBody, &bList)

	key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	w.Header().Set("Authorization", key)
	json.Unmarshal(body, &send)
	send = blacklistRework(send, bList)
	t := template.New("example")
	t = template.Must(t.ParseFiles("template/get_promo_del.html"))
	err = t.ExecuteTemplate(w, "main", send)

	if err != nil {
		log.Println(err)
	}
}

func generateInfoForm(w http.ResponseWriter) {
	var send stockSites

	key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	w.Header().Set("Authorization", key)
	token, err := utils.GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println("Token wasn't generated:", err)
	}
	send.Prod, _ = getSiteMetrics("store.ubi.com", token)
	send.Stag, _ = getSiteMetrics("store-staging.ubi.com", token)
	send.Dev, _ = getSiteMetrics("store-dev.ubi.com", token)

	t := template.New("example")
	t = template.Must(t.ParseFiles("template/get_promo_info.html"))
	err = t.ExecuteTemplate(w, "main", send)
	if err != nil {
		log.Println(err)
	}
}

func fillSend(r *http.Request, send *HostsScan) {
	r.ParseForm()
	for k, _ := range r.Form {
		switch {
		case k[:2] == "p*":
			(*send).Psites = append((*send).Psites, k[2:])
		case k[:2] == "s*":
			(*send).Ssites = append((*send).Ssites, k[2:])
		case k[:2] == "d*":
			(*send).Dsites = append((*send).Dsites, k[2:])
		case k == "prod":
			(*send).Sites = append((*send).Sites, "store.ubi.com")
		case k == "stag":
			(*send).Sites = append((*send).Sites, "store-staging.ubi.com")
		case k == "dev":
			(*send).Sites = append((*send).Sites, "store-dev.ubi.com")
		case k == "orphanedP":
			(*send).Orphaned = true
		case k == "inactiveP":
			(*send).InactiveP = true
		case k == "expiredP":
			(*send).ExpiredP = true
		case k == "inactiveC":
			(*send).InactiveC = true
		case k == "expiredC":
			(*send).ExpiredC = true
		}
	}
}

func removeIndexProm(s []promStruct, index int) []promStruct {
	return append(s[:index], s[index+1:]...)
}
func removeIndexCamp(s []campStruct, index int) []campStruct {
	return append(s[:index], s[index+1:]...)
}

func blacklistRework(hosts []hostsStruct, bList BlackList) []hostsStruct {
	for _, promo := range bList.Promotions {
		for i, host := range hosts {
			for j, site := range host.Sites {
				for k, prom := range site.PromOrphaned {
					if promo.Id == prom.PromID {
						hosts[i].Sites[j].PromOrphaned = removeIndexProm(hosts[i].Sites[j].PromOrphaned, k)
					}
				}
				for k, prom := range site.PromInactive {
					if promo.Id == prom.PromID {
						hosts[i].Sites[j].PromInactive = removeIndexProm(hosts[i].Sites[j].PromInactive, k)

					}
				}
				for k, prom := range site.PromExpired {
					if promo.Id == prom.PromID {
						hosts[i].Sites[j].PromExpired = removeIndexProm(hosts[i].Sites[j].PromExpired, k)
					}
				}
			}
		}
	}

	for _, campa := range bList.Campaigns {
		for i, host := range hosts {
			for j, site := range host.Sites {
				for k, camp := range site.CampInactive {
					if campa.Id == camp.CampID {
						hosts[i].Sites[j].CampInactive = removeIndexCamp(hosts[i].Sites[j].CampInactive, k)

					}
				}
				for k, camp := range site.CampExpired {
					if campa.Id == camp.CampID {
						hosts[i].Sites[j].CampExpired = removeIndexCamp(hosts[i].Sites[j].CampExpired, k)
					}
				}
			}
		}
	}
	return hosts
}
