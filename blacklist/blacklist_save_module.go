package blacklist

import (
	"github.com/khsadira/cleaner/auth"
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func BlacklistSaveModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {

		resp, err := auth.SendGetAuth("http://localhost:9250/blacklist/send/", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")

		if err != nil {
			log.Println("Get:", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		f, err := os.Create("./save.txt")
		if err != nil {
			log.Println("Create file:", err)
		}
		wF := bufio.NewWriter(f)
		wF.Write(body)
		wF.Flush()
		http.Redirect(w, r, "/blacklist/", 302)
	}
}