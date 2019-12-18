package example

import (
	"encoding/json"
	"github.com/khsadira/cleaner/auth"
	"github.com/khsadira/cleaner/clean/utils"
	"io/ioutil"
	"log"
	"net/http"
)

/*
**
*** +Get info form
*** +Get blacklist
*** +Send both to reworkBlacklist which rework our data with black list
*** +Send our data to Generate Form :)
**
*/

func CleanGetDataModule(w http.ResponseWriter, r *http.Request) {
	var blist Blacklist

	query := "s/-/dw/data/v19_8/sites"        // query after hostname
	endpoint := "example_get_endpoint"             // endpoint to get data | REPLACE BY OCAPI ENDPOINT TO GET DATA
	authMethod := "Bearer"                    // method used to Authorization http header | REPLACE BY THE METHOD OF OCAPI
	pathToSend := "/example/clean/deldata/"  // REPLACE BY YOUR ENDPOINT
	delEndpoint := "example_del_endpoint"          // endpoint to del data | REPLACE BY OCAPI ENDPOINT TO DEL DATA
	pathTemplate := "template/get_del.html"   //

	send := utils.GetGetData(r, query, endpoint, authMethod, pathToSend, delEndpoint)
	getBlackList(&blist)
	utils.ReworkBlacklist(&send, blist.Data)
	utils.GenerateForm(w, send, pathTemplate)
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
