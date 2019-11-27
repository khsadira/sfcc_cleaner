package cust_grp

import (
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"net/http"
)

func CleanDelDataModule(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	utils.SendFormatDelete(w, body, "customer groups", "/cstgrp/clean/del")
}

