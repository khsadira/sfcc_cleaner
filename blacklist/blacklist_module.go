package blacklist

import (
	"encoding/json"
	"github.com/khsadira/cleaner/auth"
	"io/ioutil"
	"log"
	"net/http"
)

var bList BlackList

func BlacklistModule(w http.ResponseWriter, r *http.Request) {

	if auth.TryConnection(w, r) {

		w.Write([]byte(`<html>
	<head>
		<title>UBICLEAN</title>
	</head>
	<body>
		<h1>THE CLEANER: black list</h1>
		<p>ALL BLACKLISTED ID WILL NEVER BE DELETED</p>
		<p><a href='/blacklist/show/'>show blacklisted id</a></p>
		<p><a href='/blacklist/add/'>add blacklisted id</a></p>
		<p><a href='/blacklist/del'>delete blacklisted id</a></p>
		<p><a href='/blacklist/save/'>save current blacklist</a></p>
		<p><a href='/blacklist/send/'>see black list as json</a></p>
	</body>
</html>`))
	}
}

func CreateBlackList() {
	body, err := ioutil.ReadFile("save.txt")
	if err != nil {
		log.Println("No save file")
	} else {
		json.Unmarshal(body, &bList)
	}
}
