package prom_camp

type SitesScan struct {
	Psites []string `json:"p_sites"`
	Ssites []string `json:"s_sites"`
	Dsites []string `json:"d_sites"`
}

type Scan struct {
	Total int `json:"total"`
}

type gStruct struct {
	hosts []hostsStruct `json:"hosts"`
}

type promData struct {
	Hits []hitPromStruct `json:"hits"`
}

type hitPromStruct struct {
	Id             string `json:"id"`
	Enabled        bool   `json:"enabled"`
	AssignmentInfo struct {
		CampaignId     string `json:"campaign_id"`
		ScheduleType   string `json:"schedule_type"`
		EndDate        string `json:"end_date"`
		ActiveCampaign []struct {
			CampaignId string   `json:"campaign_id"`
			Coupons    []string `json:"coupons"`
		} `json:"active_campaign_assignments"`
	} `json:"assignment_information"`
}

type stockPromStruct struct {
	promOrphaned []promStruct
	promInactive []promStruct
	promExpired  []promStruct
}

type stockCampStruct struct {
	campInactive []campStruct
	campExpired  []campStruct
}

type campData struct {
	Hits []hitCampStruct `json:"hits"`
}

type hitCampStruct struct {
	Id      string `json:"campaign_id"`
	Enabled bool   `json:"enabled"`
	EndDate string `json:"end_date"`
}
