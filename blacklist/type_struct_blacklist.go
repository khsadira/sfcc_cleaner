package blacklist

type Stock struct {
	Id     string `json:"id"`
	Encode string `json:"encode"`
}

type BlackList struct {
	Promotions []Stock `Json:"promotions"`
	Campaigns  []Stock `json:"campaigns"`
}
