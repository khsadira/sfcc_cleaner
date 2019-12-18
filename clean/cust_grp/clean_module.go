package cust_grp

import (
	"github.com/khsadira/cleaner/clean/utils"
	"net/http"
)

/*
**
*** START CLEANING BY GENERATING THE CUST_GRP CLEANING FORM
*** To get all data to clean from SFCC
*** This code is in road to be polymorphic
*** But for now you have to copy call and change variable name etc
**
 */

func CleanModule(w http.ResponseWriter, r *http.Request) {
	opt := "cust_grp"                         // to add one, separate with ',')
	pathToSend := "/customgrp/clean/getdata/" // Path used to send form on html template
	pathTemplate := "template/get_info.html"  // htlm template used to get_info

	hosts := utils.GetHosts()
	opts := utils.GetOpts(opt)
	data := utils.CreateInfo(hosts, opts, pathToSend)

	utils.GenerateForm(w, data, pathTemplate) //generate form with data to send, and html form to use
}
