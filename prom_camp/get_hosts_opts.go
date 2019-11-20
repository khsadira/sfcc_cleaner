package prom_camp

import (
	"encoding/json"
)

func getHostsOpts(body []byte) []hostsStruct {
	var scan HostsScan

	json.Unmarshal(body, &scan)
	getOptions(scan)
	ret := createHosts(scan.Sites)

	return ret
}

func createHosts(hosts []string) []hostsStruct {

	var lst []hostsStruct

	for _, host := range hosts {
		lst = append(lst, createHost(host))
	}
	return lst
}

func createHost(hostname string) hostsStruct {
	var host hostsStruct
	host.Hostname = hostname
	return host
}

func getOptions(scan HostsScan) {
	if scan.Orphaned == true {
		OrphanedP = 1
	} else {
		OrphanedP = 0
	}
	if scan.InactiveP == true {
		InactiveP = 1
	} else {
		InactiveP = 0
	}
	if scan.ExpiredP == true {
		ExpiredP = 1
	} else {
		ExpiredP = 0
	}
	if scan.InactiveC == true {
		InactiveC = 1
	} else {
		InactiveC = 0
	}
	if scan.ExpiredC == true {
		ExpiredC = 1
	} else {
		ExpiredC = 0
	}
}
