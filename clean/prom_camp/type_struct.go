package prom_camp

import "github.com/khsadira/cleaner/blacklist"

type BlackList struct {
	Promotions []blacklist.Stock `json:"promotions"`
	Campaigns  []blacklist.Stock `json:"campaigns"`
}

type HostStruct struct {
	hostname string
	check    bool
}

type HostsScan struct {
	Hosts     []HostStruct `json:"hosts"`
	Sites     []string     `json:"sites"`
	Psites    []string     `json:"p_sites"`
	Ssites    []string     `json:"s_sites"`
	Dsites    []string     `json:"d_sites"`
	Orphaned  bool         `json:"orphaned"`
	ExpiredP  bool         `json:"expired_p"`
	InactiveP bool         `json:"inactive_p"`
	InactiveC bool         `json:"inactive_c"`
	ExpiredC  bool         `json:"expired_c"`
}

type promInfos struct {
	CampID string   `json:"promotion_campaign_id"`
	CoupID []string `json:"promotion_coupon_id"`
}

type promStruct struct {
	PromID   string      `json:"promotion_id"`
	Encode   string      `json:"encode"`
	End      string      `json:"end"`
	Enabled  bool        `json:"enabled"`
	Schedule string      `json:"schedule"`
	Info     []promInfos `json:"info"`
	Check    bool        `json:"check"`
}

type campStruct struct {
	CampID string `json:"campaign_id"`
	Encode string `json:"encode"`
	End    string `json:"end"`
	Check  bool   `json:"check"`
}

type sitesStruct struct {
	SiteID string `json:"site_id"`

	PromOrphaned []promStruct `json:"promotion_orphaned"`
	PromInactive []promStruct `json:"promotion_inactived"`
	PromExpired  []promStruct `json:"promotion_expired"`
	CampInactive []campStruct `json:"campaign_inactived"`
	CampExpired  []campStruct `json:"campaign_expired"`
}

type hostsStruct struct {
	Hostname string        `json:"hostname"`
	Sites    []sitesStruct `json:"sites"`
}

type Sites struct {
	Data []struct {
		Type string `json:"_type"`
		ID   string `json:"id"`
	} `json:"data"`
	Total int `json:"total"`
	Count int `json:"count"`
}

type stockSites struct {
	Prod []string
	Stag []string
	Dev  []string
}

type Token struct {
	AccessToken string `json:"access_token"`
}
