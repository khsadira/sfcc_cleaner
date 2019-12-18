package example

import (
	"github.com/khsadira/cleaner/clean/utils"
	"net/http"
)

/*
**
*** THIS IS AN EXAMPLE TO NOW HOW YOU CAN IMPORT EASILY NEW OBJECT
***
*** COPY PAST THIS FOLDER (EXAMPLE) AND RENAME IT BY THE NAME THAT YOU WANT(by example configuration_slot), AS YOU CAN SEE IN CUST_GRP FOLDER.
*** IN FIRST, YOU MUST CHANGE ALL PACKAGE NAME ON THE TOP OF ANY FILES IN THE FOLDER
*** AFTER YOU HAVE TO GO IN EACH FILES OF THIS FOLDER AND CHANGE EVERY VARIABLE NEEDED TO BE CHANGE (it's explicitly written)
***
*** NOW, YOU HAVE TO GO IN THE MAIN.GO FILE (in the main folder)
*** IN THE TOP OF THE FILE, GO IN IMPORT AND ADD THE PACKAGE NAME THAT YOU CHOSE BEFORE(you can use as example the example import)
*** GO IN "homePage" AND ADD YOUR OWN <a href> USING YOUR OWN ENDPOINT (you only have to specified it here)
*** FINALLY YOU HAVE TO ADD THE 3 NEEDED ENDPOINT, ON THE ROUTER FUNCTION, SEE AS EXAMPLE, EXAMPLE ENDPOINT
***
*** IF YOU WANT YOU CAN ADD FILE THAT YOU CREATE ON THE MAKEFILE (only for format support)
**
*/

/*
**
*** START CLEANING BY GENERATING THE EXAMPLE CLEANING FORM
*** To get all data to clean from SFCC
*** This code is in road to be polymorphic
*** But for now you have to copy call and change variable name etc
**
*/

func CleanModule(w http.ResponseWriter, r *http.Request) {
	opt := "example_opt"                         // to add one, separate with ',') | REPLACE BY THE NAME OF YOUR OPTIONS
	pathToSend := "/example/clean/getdata/" // Path used to send form on html template | REPLACE EXAMPLE BY YOUR ENDPOINT
	pathTemplate := "template/get_info.html"  // htlm template used to get_info

	hosts := utils.GetHosts()
	opts := utils.GetOpts(opt)
	data := utils.CreateInfo(hosts, opts, pathToSend)

	utils.GenerateForm(w, data, pathTemplate) //generate form with data to send, and html form to use
}
