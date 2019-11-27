package utils

type SitesScan struct {
	Psites []string `json:"p_sites"`
	Ssites []string `json:"s_sites"`
	Dsites []string `json:"d_sites"`
}

type Sites struct {
	Data []struct {
		ID   string `json:"id"`
	} `json:"data"`
	Total int `json:"total"`
	Count int `json:"count"`
}

type Host struct {
	Inst  string
	ID    string
	Name  string
	Host  string
	Sites []string
}

type Opts struct {
	ID string
	Name string
}

type InfoData struct {
	Path string
	Env  []Host
	Opts []Opts
}

type DataStruct struct {
	ID string `json:"id"`
	Encoded string `json:"encoded"`
}

type OptsStruct struct {
	Name string
	Data []DataStruct
}

type SitesStruct struct {
	SitesID string
	Opts []OptsStruct
	Data []DataStruct
}

type HostStruct struct {
	Instance string
	Hostname string
	Sites []SitesStruct
}

type Global struct {
	Options []string
	Path string
	Hosts []HostStruct
}