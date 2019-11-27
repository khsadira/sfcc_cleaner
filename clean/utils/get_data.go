package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetGetData(r *http.Request, query string, endpoint string, auth string, path string) Global {
	hosts, opts := FormatInfo(r)

	for i, host := range hosts {
		for j, site := range host.Sites {
			nbThread := len(site.Opts)
			g := make(chan bool, nbThread)
			for k, _ := range site.Opts {
				go func(g chan bool, i, j, k int) {
					query := fmt.Sprintf("https://%s/%s/%s/%s", host.Hostname, query, site.SitesID, endpoint)
					datas := GetDataSFCC(query, auth)
					hosts[i].Sites[j].Opts[k].Data = SwapData(datas)
					g <- true
				}(g, i, j, k)
			}
			for e := 0; e < nbThread; e++ {
				<-g
			}
		}
	}
	return Global{opts, path, hosts}
}

func SwapData(datas []interface{}) []DataStruct {
	var ret []DataStruct

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
	}
	return ret
}


func FormatInfo(r *http.Request) ([]HostStruct, []string) {
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


	if len(opts) > 0 {
		for _, opt := range opts {
			for i, _ := range prod.Sites {
				prod.Sites[i].Opts = append(prod.Sites[i].Opts, OptsStruct{Name: opt})
			}
			for i, _ := range stag.Sites {
				stag.Sites[i].Opts = append(stag.Sites[i].Opts, OptsStruct{Name: opt})
			}
			for i, _ := range dev.Sites {
				dev.Sites[i].Opts = append(dev.Sites[i].Opts, OptsStruct{Name: opt})
			}
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
	return ret, opts
}