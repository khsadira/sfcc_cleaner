package cust_grp

import "github.com/khsadira/cleaner/blacklist"

type Blacklist struct {
	Data []blacklist.Stock `json:"customer_groups"`
}