package blacklist

type Stock struct {
	Id     string `json:"id"`
	Encode string `json:"encode"`
}

type BlackList struct {
	Promotions []Stock `json:"promotions"`
	Campaigns  []Stock `json:"campaigns"`
	CustGrps   []Stock `json:"customer_groups"`
}
