package prom_camp

func addJsonToStruct(site *sitesStruct, promData promData, campData campData) {
	var promDataStock stockPromStruct
	var campDataStock stockCampStruct

	g := make(chan bool, 2)
	go promDataToStruct(promData, &promDataStock, g)
	go campDataToStruct(campData, &campDataStock, g)
	for i := 0; i < 2; i++ {
		<-g
	}
	addDataToSite(site, promDataStock, campDataStock)
}

func promDataToStruct(data promData, ret *stockPromStruct, g chan bool) {
	for _, hit := range data.Hits {
		fillPromAttribut(ret, hit)
	}
	g <- true
}

func campDataToStruct(data campData, ret *stockCampStruct, g chan bool) {
	for _, hit := range data.Hits {
		fillCampAttribut(ret, hit)
	}
	g <- true
}

func addDataToSite(site *sitesStruct, promData stockPromStruct, campData stockCampStruct) {
	ch := OrphanedP + InactiveP + ExpiredP + InactiveC + ExpiredC
	h := make(chan bool, ch)
	if OrphanedP == 1 {
		go func(h chan bool) {
			(*site).PromOrphaned = append((*site).PromOrphaned, promData.promOrphaned...)
			h <- true
		}(h)

	}
	if InactiveP == 1 {
		go func(h chan bool) {
			(*site).PromInactive = append((*site).PromInactive, promData.promInactive...)
			h <- true
		}(h)
	}
	if ExpiredP == 1 {
		go func(h chan bool) {
			(*site).PromExpired = append((*site).PromExpired, promData.promExpired...)
			h <- true
		}(h)
	}
	if InactiveC == 1 {
		go func(h chan bool) {
			(*site).CampInactive = append((*site).CampInactive, campData.campInactive...)
			h <- true
		}(h)
	}
	if ExpiredC == 1 {
		go func(h chan bool) {
			(*site).CampExpired = append((*site).CampExpired, campData.campExpired...)
			h <- true
		}(h)
	}

	for i := 0; i < ch; i++ {
		<-h
	}
}
