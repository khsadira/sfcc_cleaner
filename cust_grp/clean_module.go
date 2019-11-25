package cust_grp

import (
	"github.com/khsadira/cleaner/auth"
	"github.com/khsadira/cleaner/utils"
	"net/http"
)

func CleanModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {
		hosts := utils.GetHosts()

		opts := utils.GetOpts("cust_grp")
		send := utils.CreateInfo(hosts, opts, "/customgrp/clean/getdata/")

		utils.GenerateForm(w, send, "template/get_info.html")
	}
}