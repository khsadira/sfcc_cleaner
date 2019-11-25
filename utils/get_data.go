package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetGetData(r *http.Request, query string, endpoint string, auth string, path string) Global {
	hosts := FormatInfo(r)

	for i, host := range hosts {
		nbThread := len(host.Sites)
		g := make(chan bool, nbThread)
		for j, site := range host.Sites {
			println(site.SitesID)
			go func(g chan bool, i, j int) {
				query := fmt.Sprintf("https://%s/%s/%s/%s", host.Hostname, query, site.SitesID, endpoint)
				datas := GetDataSFCC(query, auth)
				hosts[i].Sites[j].Data = SwapData(datas)
				g <- true
			}(g, i, j)
		}
		for e := 0; e < nbThread; e++ {
			<-g
		}
	}
	return Global{path, hosts}
}

func SwapData(datas []interface{}) []DataStruct {
	var ret []DataStruct

	var ids []string
	for _, data := range datas {
		m, ok := data.(map[string]interface{})
		if !ok {
			log.Printf("want type map[string]interface: got %T\n", data)
			return nil
		}
		for k, v := range m {
			if k == "data" {
				vstring := fmt.Sprintf("%v", v)
				splits := strings.Split(vstring, "map")
				for _, split := range splits {
					indexId := strings.Index(split, " id:")
					indexLink := strings.Index(split, " link:")
					if indexId >= 0 && indexLink >= 0 {
						id := split[indexId+4:indexLink]
						ret = append(ret, DataStruct{id, hex.EncodeToString([]byte(id))})
					}
				}
			}
		}
		for _, id := range ids {
			println(id)
		}
	}
	return ret
}


func FormatInfo(r *http.Request) []HostStruct {
	var prod, stag, dev HostStruct
	var p, s, d bool
	var opts []string

	r.ParseForm()
	for k := range r.Form {
		switch {
		case k[:2] == "p*":
			prod.Sites = append(prod.Sites, SitesStruct{SitesID: k[2:]})
		case k[:2] == "s*":
			stag.Sites = append(stag.Sites, SitesStruct{SitesID: k[2:]})
		case k[:2] == "d*":
			dev.Sites = append(dev.Sites, SitesStruct{SitesID: k[2:]})
		case k == "prod":
			p = true
		case k == "stag":
			s = true
		case k == "dev":
			d = true
		case k[:4] == "opt*":
			opts = append(opts, k[4:])
		}
	}

	var ret []HostStruct

	if p {
		prod.Hostname = "store.ubi.com"
		prod.Instance = "p"
		ret = append(ret, prod)
	}
	if s {
		stag.Hostname = "store-staging.ubi.com"
		stag.Instance = "s"
		ret = append(ret, stag)
	}
	if d {
		dev.Hostname = "store-dev.ubi.com"
		dev.Instance = "d"
		ret = append(ret, dev)
	}
	return ret
}