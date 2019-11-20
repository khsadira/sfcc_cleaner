package prom_camp

func fillMetric(host string, site *sitesStruct, token string, c chan bool) {
	var promData promData
	var campData campData

	e := make(chan bool, 2)
	go collectPromJSON(&promData, host, site.SiteID, token, e)
	go collectCampJSON(&campData, host, site.SiteID, token, e)
	for i := 0; i < 2; i++ {
		<-e
	}
	addJsonToStruct(site, promData, campData)
	c <- true
}

func collectPromJSON(data *promData, host string, site string, token string, e chan bool) {
	searchs := []string{"true", "false"}
	*data = getPromoJSONs(host, site, token, searchs)

	e <- true
}

func collectCampJSON(data *campData, host string, site string, token string, e chan bool) {
	searchs := []string{"true", "false"}
	*data = getCampJSONs(host, site, token, searchs)
	e <- true
}
