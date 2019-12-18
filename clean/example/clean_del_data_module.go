package example

import (
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"net/http"
)

/*
**
*** Sending data selected by users to confirmation page
**
*/

func CleanDelDataModule(w http.ResponseWriter, r *http.Request) {
	dataName := "example data" // REPLACE BY YOUR DATA NAME
	pathToDelete := "/clean/delete/"

	body, _ := ioutil.ReadAll(r.Body)
	utils.SendConfirmDelete(w, body, dataName, pathToDelete)
}
