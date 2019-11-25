package cust_grp

import (
	"github.com/khsadira/cleaner/utils"
	"net/http"
)

/*
*** To get all data to clean from SFCC
*** This code is in road to be polymorphic
*** But for now you have to implement by yourself all feature
 */

func CleanGetDataModule(w http.ResponseWriter, r *http.Request) {
	query := "s/-/dw/data/v19_8/sites"
	endpoint := "customer_groups"
	auth := "Bearer"

	send := utils.GetGetData(r, query, endpoint, auth, "/customgrp/clean/deldata/")
	utils.GenerateForm(w, send, "template/get_del.html")
}


