package cust_grp

import "github.com/khsadira/cleaner/blacklist"

type BlackList struct {
	CustGrps []blacklist.Stock `json:"customer_groups"`
}
