package test_cleaner

import (
	"github.com/khsadira/cleaner/clean/utils"
	"testing"
)

func TestGetHosts(t *testing.T) {
	hosts := utils.GetHosts()
	if len(hosts) == 0 {
		t.Errorf("GetHosts was incorrrect, got no slice")
	}
	for _, host := range hosts {
		if len(host) == 0 || host == "" {
			t.Errorf("GetHosts was incorrect, got nothing in one string of the slice")
		}
	}
}

func TestGetSites(t *testing.T) {
	sites := utils.GetSites("store.ubi.com")
	if len(sites) == 0 {
		t.Errorf("GetSites was incorrect, got no slice")
	}
	for _, site := range sites {
		if len(site) == 0 || site == "" {
			t.Errorf("GetSites was incorrect, got nothing in one string of the slice")
		}
	}
	sites = utils.GetSites("")
}

func TestCreateInfo(t *testing.T) {
	hosts := []string{"store.ubi.com", "store-staging.ubi.com", "store-dev.ubi.com"}
	opts := []string{"ExpiredC", "Orphaned"}
	path := "/clean/"

	data := utils.CreateInfo(hosts, opts, path)

	if len(data.Env) != 3 {
		t.Errorf("CreateInfo was incorrect, we didn't get 3 instances")
	}
	if len(data.Opts) != 2 {
		t.Errorf("CreateInfo was incorrect, we didn't get 2 opts")
	}
	data = utils.CreateInfo(hosts, nil, "")
	data = utils.CreateInfo(nil, opts, "")
	data = utils.CreateInfo(nil, nil, "")
}
