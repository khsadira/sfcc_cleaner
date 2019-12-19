package prom_camp

import (
	"encoding/hex"
	"time"
)

func fillPromAttribut(ret *stockPromStruct, hit hitPromStruct) {
	resO, resI, resE := findPromType(hit)
	if OrphanedP == 1 && resO == 1 {
		(*ret).promOrphaned = append((*ret).promOrphaned, createPromStruct(hit))
	} else if InactiveP == 1 && resI == 1 {
		(*ret).promInactive = append((*ret).promInactive, createPromStruct(hit))
	} else if ExpiredP == 1 && resE == 1 {
		(*ret).promExpired = append((*ret).promExpired, createPromStruct(hit))
	}
}

func createPromStruct(hit hitPromStruct) promStruct {
	var ret promStruct
	var infos []promInfos

	ret.PromID = hit.Id
	ret.Encode = hex.EncodeToString([]byte(ret.PromID))
	ret.Enabled = hit.Enabled
	ret.End = hit.AssignmentInfo.EndDate
	ret.Schedule = hit.AssignmentInfo.ScheduleType

	for _, data := range hit.AssignmentInfo.ActiveCampaign {
		var info promInfos

		info.CoupID = data.Coupons
		info.CampID = data.CampaignId
		infos = append(infos, info)
	}

	ret.Check = false
	ret.Info = infos

	return ret
}

func findPromType(hit hitPromStruct) (int, int, int) {
	resO := 0
	resI := 0
	resE := 0

	if hit.Enabled == false {
		resI = 1
	}
	if hit.AssignmentInfo.ScheduleType == "none" {
		resO = 1
	}

	end, _ := time.Parse(time.RFC3339, hit.AssignmentInfo.EndDate)

	println("date:", end.Unix())

	if end.Unix() <= date && end.Unix() >= 0 {
		resE = 1
	}
	return resO, resI, resE
}
