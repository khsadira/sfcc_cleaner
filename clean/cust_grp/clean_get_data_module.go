package cust_grp

import (
	"encoding/json"
	"github.com/khsadira/cleaner/auth"
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"log"
	"net/http"
)

/*
*** To get all data to clean from SFCC
*** This code is in road to be polymorphic
*** But for now you have to implement by yourself all feature
 */

func CleanGetDataModule(w http.ResponseWriter, r *http.Request) {
	var blist Blacklist

	query := "s/-/dw/data/v19_8/sites"
	endpoint := "customer_groups"
	auth := "Bearer"

	send := utils.GetGetData(r, query, endpoint, auth, "/customgrp/clean/deldata/")
	getBlackList(&blist)
	utils.ReworkBlacklist(&send, blist.Data)
	utils.GenerateForm(w, send, "template/get_del.html")
}

func getBlackList(blist *Blacklist) {
	resp, err := auth.SendGetAuth("http://localhost:9250/blacklist/send/", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	blBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(blBody, blist)
}
