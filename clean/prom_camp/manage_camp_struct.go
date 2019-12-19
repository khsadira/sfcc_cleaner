package prom_camp

import (
	"encoding/hex"
	"time"
)

func fillCampAttribut(ret *stockCampStruct, hit hitCampStruct) {
	resI, resE := findCampType(hit)
	if InactiveC == 1 && resI == 1 {
		(*ret).campInactive = append((*ret).campInactive, createCampStruct(hit))
	} else if ExpiredC == 1 && resE == 1 {
		(*ret).campExpired = append((*ret).campExpired, createCampStruct(hit))
	}
}

func createCampStruct(hit hitCampStruct) campStruct {
	var ret campStruct

	ret.CampID = hit.Id
	ret.Encode = hex.EncodeToString([]byte(ret.CampID))
	ret.End = hit.EndDate
	ret.Check = false

	return ret
}

func findCampType(hit hitCampStruct) (int, int) {
	var resI, resE int

	if hit.Enabled == false {
		resI = 1
	}

	end, _ := time.Parse(time.RFC3339, hit.EndDate)
	if end.Unix() <= date && end.Unix() >= 0 {
		resE = 1
	}
	return resI, resE
}
