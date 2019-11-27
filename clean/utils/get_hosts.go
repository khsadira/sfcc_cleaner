package utils

func GetHosts() []string {
	/*	prod := os.Getenv("ENV_PROD_SFCC")
		stag := os.Getenv("ENV_STAG_SFCC")
		dev := os.Getenv("ENV_DEV_SFCC")
	*/
	prod := "store.ubi.com"
	stag := "store-staging.ubi.com"
	dev := "store-dev.ubi.com"
	return []string{prod, stag, dev}
}
