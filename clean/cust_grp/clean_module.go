package cust_grp

import (
	"github.com/khsadira/cleaner/clean/utils"
	"net/http"
)

/*
*** START CLEANING BY GENERATING THE CUST_GRP CLEANING FORM
*** To get all data to clean from SFCC
*** This code is in road to be polymorphic
*** But for now you have to copy call and change variable name etc
 */

func CleanModule(w http.ResponseWriter, r *http.Request) {
	hosts := utils.GetHosts()

	opts := utils.GetOpts("cust_grp")
	send := utils.CreateInfo(hosts, opts, "/customgrp/clean/getdata/")

	utils.GenerateForm(w, send, "template/get_info.html")
}
