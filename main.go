package main

import (
	"github.com/khsadira/cleaner/auth"
	"github.com/khsadira/cleaner/blacklist"
	"github.com/khsadira/cleaner/prom_camp"
	"log"
	"net/http"
)

func routers() *http.ServeMux {
	mux := http.NewServeMux()

	auth.AuthStart(mux, "/promo/clean/", prom_camp.CleanModule)
	auth.AuthStart(mux, "/promo/clean/del", prom_camp.CleanDelModule)
	auth.AuthStart(mux, "/promo/script/info/", prom_camp.ScriptInfoModule)
	auth.AuthStart(mux, "/promo/script/delete/", prom_camp.ScriptDeleteModule)

	auth.AuthStart(mux, "/blacklist/", blacklist.BlacklistModule)
	auth.AuthStart(mux, "/blacklist/show/", blacklist.BlacklistShowModule)
	auth.AuthStart(mux, "/blacklist/add/", blacklist.BlacklistAddModule)
	auth.AuthStart(mux, "/blacklist/add/values/", blacklist.BlacklistAddValuesModule)
	auth.AuthStart(mux, "/blacklist/send/", blacklist.BlacklistSendModule)
	auth.AuthStart(mux, "/blacklist/del/", blacklist.BlacklistDelModule)
	auth.AuthStart(mux, "/blacklist/del/values/", blacklist.BlacklistDelValuesModule)
	auth.AuthStart(mux, "/blacklist/save/", blacklist.BlacklistSaveModule)

	mux.Handle("/clean/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>UBICLEAN</title></head>
			<body>
			<h1>THE CLEANER</h1>
			<p><a href='/promo/clean/'>start cleaner for promotions and campaigns</a></p>
			<p><a href='/customgrp/clean/'>start cleaner for customer groups</a></p>
			</body>
			</html>`))
	}))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>UBICLEAN</title></head>
			<body>
			<h1>THE CLEANER</h1>
			<p><a href='/clean/'>start clean</a></p>
			<p><a href='/blacklist/'>manage blacklist</a></p>
			</body>
			</html>`))
	}))
	log.Fatal(http.ListenAndServe(":9250", mux))

	return mux
}

func main() {
	blacklist.CreateBlackList()
	if r := routers(); r != nil {
		log.Fatal("Server exited:", http.ListenAndServe(":9250", r))
	}
}
