package prom_camp

func makeSites(host string, scan SitesScan) ([]sitesStruct, int) {
	var sites []sitesStruct
	var size int

	if host == "store.ubi.com" {
		for _, site := range scan.Psites {
			sites = append(sites, sitesStruct{SiteID: site})
			size++
		}
	} else if host == "store-staging.ubi.com" {
		for _, site := range scan.Ssites {
			sites = append(sites, sitesStruct{SiteID: site})
			size++
		}
	} else if host == "store-dev.ubi.com" {
		for _, site := range scan.Dsites {
			sites = append(sites, sitesStruct{SiteID: site})
			size++
		}
	}
	return sites, size
}
