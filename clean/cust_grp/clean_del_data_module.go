package cust_grp

import (
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"net/http"
)

/*
*** Sending data selected by users to confirmation page
 */

func CleanDelDataModule(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	utils.SendConfirmDelete(w, body, "customer groups", "/clean/delete/")
}
