package prom_camp

import (
	"encoding/json"
	"github.com/khsadira/cleaner/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	OrphanedP int
	InactiveP int
	ExpiredP  int
	InactiveC int
	ExpiredC  int
	date      int64
)

func ScriptInfoModule(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var lst gStruct

		body, _ := ioutil.ReadAll(r.Body)
		lst.hosts = getHostsOpts(body)

		token, err := utils.GetToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
		if err != nil {
			log.Println("token can't be reach", err)
			return
		}

		var scanSites SitesScan
		json.Unmarshal(body, &scanSites)

		date = time.Now().Unix() - (7 * 24 * 60)
		for i := range lst.hosts {
			var size int
			lst.hosts[i].Sites, size = makeSites(lst.hosts[i].Hostname, scanSites)
			c := make(chan bool, size)
			for j := range lst.hosts[i].Sites {
				fillMetric(lst.hosts[i].Hostname, &lst.hosts[i].Sites[j], token, c)
			}
			for i := 0; i < size; i++ {
				<-c
			}
		}
		OrphanedP, InactiveP, InactiveC, ExpiredP, ExpiredC, date = 0, 0, 0, 0, 0, 0
		ret := byteMetrics(lst)
		w.Write(ret)
	}
}

func byteMetrics(hosts gStruct) []byte {
	ret, err := json.Marshal(hosts.hosts)
	if err != nil {
		log.Println(err)
	}
	return ret
}
